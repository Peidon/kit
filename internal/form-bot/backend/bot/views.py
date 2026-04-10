import json
import uuid
from datetime import datetime

from django.db import transaction
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods

from .models import Education, Experience, UserProfile


DATE_FORMAT = "%d-%m-%Y"
PROFILE_REQUIRED_FIELDS = [
    "firstname",
    "lastname",
    "email",
    "phone",
    "location",
    "city",
    "country",
    "linkedin",
]
EDUCATION_REQUIRED_FIELDS = [
    "school",
    "field",
    "degree",
    "startDate",
    "endDate",
]
EXPERIENCE_REQUIRED_FIELDS = [
    "company",
    "title",
    "startDate",
    "endDate",
    "description",
]


def _missing_fields(data, required_fields):
    return [field for field in required_fields if not data.get(field)]


def _parse_date(value, field_name):
    try:
        return datetime.strptime(value, DATE_FORMAT).date()
    except (TypeError, ValueError) as exc:
        raise ValueError(f"{field_name} must use DD-MM-YYYY format") from exc


def _format_date(value):
    if value is None:
        return None
    return value.strftime(DATE_FORMAT)


def _serialize_user_info(profile):
    return {
        "profile": {
            "firstname": profile.firstname,
            "lastname": profile.lastname,
            "email": profile.email,
            "phone": profile.phone,
            "location": profile.location,
            "city": profile.city,
            "country": profile.country,
            "linkedin": profile.linkedin,
        },
        "educations": [
            {
                "school": education.school,
                "field": education.field,
                "degree": education.degree,
                "startDate": _format_date(education.start_date),
                "endDate": _format_date(education.end_date),
            }
            for education in profile.educations.all().order_by("created_at")
        ],
        "experiences": [
            {
                "company": experience.company,
                "title": experience.title,
                "startDate": _format_date(experience.start_date),
                "endDate": _format_date(experience.end_date),
                "description": experience.description,
            }
            for experience in profile.experiences.all().order_by(
                "created_at", "company"
            )
        ],
    }


@csrf_exempt
@require_http_methods(["GET", "POST"])
def user_info(request):
    if request.method == "GET":
        user_id = request.GET.get("user_id")
        if not user_id:
            return JsonResponse(
                {"error": "user_id query parameter is required."},
                status=400,
            )

        try:
            parsed_user_id = uuid.UUID(user_id)
        except ValueError:
            return JsonResponse({"error": "user_id must be a valid UUID."}, status=400)

        profile = (
            UserProfile.objects.prefetch_related("educations", "experiences")
            .filter(id=parsed_user_id)
            .first()
        )
        if profile is None:
            return JsonResponse({"error": "User not found."}, status=404)

        return JsonResponse(_serialize_user_info(profile), status=200)

    try:
        payload = json.loads(request.body)
    except json.JSONDecodeError:
        return JsonResponse({"error": "Request body must be valid JSON."}, status=400)

    profile_data = payload.get("profile")
    educations_data = payload.get("educations", [])
    experiences_data = payload.get("experiences", [])

    if not isinstance(profile_data, dict):
        return JsonResponse({"error": "profile must be an object."}, status=400)
    if not isinstance(educations_data, list):
        return JsonResponse({"error": "educations must be an array."}, status=400)
    if not isinstance(experiences_data, list):
        return JsonResponse({"error": "experiences must be an array."}, status=400)

    missing_profile_fields = _missing_fields(profile_data, PROFILE_REQUIRED_FIELDS)
    if missing_profile_fields:
        return JsonResponse(
            {
                "error": "profile is missing required fields.",
                "missingFields": missing_profile_fields,
            },
            status=400,
        )

    for index, education in enumerate(educations_data):
        if not isinstance(education, dict):
            return JsonResponse(
                {"error": f"educations[{index}] must be an object."},
                status=400,
            )
        missing_fields = _missing_fields(education, EDUCATION_REQUIRED_FIELDS)
        if missing_fields:
            return JsonResponse(
                {
                    "error": f"educations[{index}] is missing required fields.",
                    "missingFields": missing_fields,
                },
                status=400,
            )

    for index, experience in enumerate(experiences_data):
        if not isinstance(experience, dict):
            return JsonResponse(
                {"error": f"experiences[{index}] must be an object."},
                status=400,
            )
        missing_fields = _missing_fields(experience, EXPERIENCE_REQUIRED_FIELDS)
        if missing_fields:
            return JsonResponse(
                {
                    "error": f"experiences[{index}] is missing required fields.",
                    "missingFields": missing_fields,
                },
                status=400,
            )

    try:
        with transaction.atomic():
            profile, _created = UserProfile.objects.update_or_create(
                email=profile_data["email"],
                defaults={
                    "firstname": profile_data["firstname"],
                    "lastname": profile_data["lastname"],
                    "phone": profile_data["phone"],
                    "location": profile_data["location"],
                    "city": profile_data["city"],
                    "country": profile_data["country"],
                    "linkedin": profile_data["linkedin"],
                },
            )

            Education.objects.filter(profile=profile).delete()
            Experience.objects.filter(profile=profile).delete()

            Education.objects.bulk_create(
                [
                    Education(
                        profile=profile,
                        school=education["school"],
                        field=education["field"],
                        degree=education["degree"],
                        start_date=_parse_date(education["startDate"], "startDate"),
                        end_date=_parse_date(education["endDate"], "endDate"),
                    )
                    for education in educations_data
                ]
            )
            Experience.objects.bulk_create(
                [
                    Experience(
                        profile=profile,
                        company=experience["company"],
                        title=experience["title"],
                        start_date=_parse_date(experience["startDate"], "startDate"),
                        end_date=_parse_date(experience["endDate"], "endDate"),
                        description=experience["description"],
                    )
                    for experience in experiences_data
                ]
            )
    except ValueError as exc:
        return JsonResponse({"error": str(exc)}, status=400)

    return JsonResponse(
        {
            "message": "User info saved successfully.",
            "userId": str(profile.id),
            "educationCount": len(educations_data),
            "experienceCount": len(experiences_data),
        },
        status=201,
    )

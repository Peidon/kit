import json
import uuid

from django.test import Client, TestCase

from .models import Education, Experience, UserProfile


class UserInfoViewTests(TestCase):
    def setUp(self):
        self.client = Client()
        self.payload = {
            "profile": {
                "firstname": "Avery",
                "lastname": "Chen",
                "email": "avery.chen@example.com",
                "phone": "0400000000",
                "location": "123 Example Street, Testville, NSW 2000",
                "city": "Testville",
                "country": "Australia",
                "linkedin": "https://www.linkedin.com/in/avery-chen-example/",
            },
            "educations": [
                {
                    "school": "Example University",
                    "field": "Computer Science",
                    "degree": "Master",
                    "endDate": "10-06-2020",
                    "startDate": "10-09-2017",
                },
                {
                    "school": "Sample Institute of Technology",
                    "field": "Information Systems",
                    "degree": "Bachelor",
                    "endDate": "10-06-2017",
                    "startDate": "01-09-2013",
                },
            ],
            "experiences": [
                {
                    "company": "Example Tech Pty Ltd",
                    "title": "Backend Engineer",
                    "startDate": "17-06-2020",
                    "endDate": "18-06-2025",
                    "description": "Built internal APIs and data processing services.",
                }
            ],
        }

    def test_post_user_info_saves_nested_records(self):
        response = self.client.post(
            "/user_info",
            data=json.dumps(self.payload),
            content_type="application/json",
        )

        self.assertEqual(response.status_code, 201)
        self.assertEqual(UserProfile.objects.count(), 1)
        self.assertEqual(Education.objects.count(), 2)
        self.assertEqual(Experience.objects.count(), 1)

        profile = UserProfile.objects.get(email="avery.chen@example.com")
        self.assertEqual(profile.firstname, "Avery")
        self.assertEqual(response.json()["userId"], str(profile.id))

    def test_post_user_info_replaces_existing_nested_records(self):
        self.client.post(
            "/user_info",
            data=json.dumps(self.payload),
            content_type="application/json",
        )

        updated_payload = json.loads(json.dumps(self.payload))
        updated_payload["profile"]["phone"] = "0411111111"
        updated_payload["educations"] = updated_payload["educations"][:1]
        updated_payload["experiences"][0]["title"] = "Senior Backend Engineer"

        response = self.client.post(
            "/user_info",
            data=json.dumps(updated_payload),
            content_type="application/json",
        )

        self.assertEqual(response.status_code, 201)
        self.assertEqual(UserProfile.objects.count(), 1)
        self.assertEqual(Education.objects.count(), 1)
        self.assertEqual(Experience.objects.count(), 1)
        self.assertEqual(
            UserProfile.objects.get(email="avery.chen@example.com").phone,
            "0411111111",
        )

    def test_post_user_info_rejects_invalid_date_format(self):
        self.payload["educations"][0]["startDate"] = "2017-09-10"

        response = self.client.post(
            "/user_info",
            data=json.dumps(self.payload),
            content_type="application/json",
        )

        self.assertEqual(response.status_code, 400)
        self.assertIn("DD-MM-YYYY", response.json()["error"])

    def test_get_user_info_returns_nested_payload(self):
        self.client.post(
            "/user_info",
            data=json.dumps(self.payload),
            content_type="application/json",
        )
        profile = UserProfile.objects.get(email="avery.chen@example.com")

        response = self.client.get("/user_info", {"user_id": str(profile.id)})

        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json(), self.payload)

    def test_get_user_info_requires_valid_uuid(self):
        response = self.client.get("/user_info", {"user_id": "not-a-uuid"})

        self.assertEqual(response.status_code, 400)
        self.assertEqual(response.json()["error"], "user_id must be a valid UUID.")

    def test_get_user_info_returns_not_found_for_unknown_uuid(self):
        response = self.client.get("/user_info", {"user_id": str(uuid.uuid4())})

        self.assertEqual(response.status_code, 404)
        self.assertEqual(response.json()["error"], "User not found.")

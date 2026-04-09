import json

from django.test import Client, TestCase

from .models import Education, Experience, UserProfile


class UserInfoViewTests(TestCase):
    def setUp(self):
        self.client = Client()
        self.payload = {
            "profile": {
                "firstname": "Peidong",
                "lastname": "Xu",
                "email": "peidong.xu94@gmail.com",
                "phone": "0417909812",
                "location": "11 Hillcrest Ave, Epping, NSW 2121",
                "city": "Sydney",
                "country": "Australia",
                "linkedin": "https://www.linkedin.com/in/peidong-xu-9aab4414b/",
            },
            "educations": [
                {
                    "school": "Nanjing University",
                    "field": "Software Engineering",
                    "degree": "Master",
                    "endDate": "10-06-2020",
                    "startDate": "10-09-2017",
                },
                {
                    "school": "East China Institute of Geology",
                    "field": "Embedded Systems",
                    "degree": "Bachelor",
                    "endDate": "10-06-2017",
                    "startDate": "01-09-2013",
                },
            ],
            "experiences": [
                {
                    "company": "Shopee",
                    "title": "Software Engineer",
                    "startDate": "17-06-2020",
                    "endDate": "18-06-2025",
                    "description": "I mainly worked on the Chatbot platform.",
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

        profile = UserProfile.objects.get(email="peidong.xu94@gmail.com")
        self.assertEqual(profile.firstname, "Peidong")

    def test_post_user_info_replaces_existing_nested_records(self):
        self.client.post(
            "/user_info",
            data=json.dumps(self.payload),
            content_type="application/json",
        )

        updated_payload = json.loads(json.dumps(self.payload))
        updated_payload["profile"]["phone"] = "0400000000"
        updated_payload["educations"] = updated_payload["educations"][:1]
        updated_payload["experiences"][0]["title"] = "Senior Software Engineer"

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
            UserProfile.objects.get(email="peidong.xu94@gmail.com").phone,
            "0400000000",
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

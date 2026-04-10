import uuid

from django.db import models


class UserProfile(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    firstname = models.CharField(max_length=100)
    lastname = models.CharField(max_length=100)
    email = models.EmailField(unique=True)
    phone = models.CharField(max_length=32, blank=True, null=True)
    location = models.CharField(max_length=255, blank=True, null=True)
    city = models.CharField(max_length=100, blank=True, null=True)
    country = models.CharField(max_length=100, blank=True, null=True)
    linkedin = models.URLField(blank=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = "user_profile"

    def __str__(self) -> str:
        return f"{self.firstname} {self.lastname} <{self.email}>"


class Education(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    profile = models.ForeignKey(
        UserProfile,
        related_name="educations",
        on_delete=models.CASCADE,
        db_column="user_id",
    )
    school = models.CharField(max_length=255)
    field = models.CharField(max_length=255, blank=True, null=True)
    degree = models.CharField(max_length=100, blank=True, null=True)
    start_date = models.DateField(blank=True, null=True)
    end_date = models.DateField(blank=True, null=True)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        db_table = "user_education"

    def __str__(self) -> str:
        return f"{self.school} - {self.degree}"


class Experience(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    profile = models.ForeignKey(
        UserProfile,
        related_name="experiences",
        on_delete=models.CASCADE,
        db_column="user_id",
    )
    company = models.CharField(max_length=255)
    title = models.CharField(max_length=255, blank=True, null=True)
    start_date = models.DateField(blank=True, null=True)
    end_date = models.DateField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        db_table = "user_experience"

    def __str__(self) -> str:
        return f"{self.company} - {self.title}"

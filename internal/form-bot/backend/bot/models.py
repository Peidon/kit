from django.db import models


class UserProfile(models.Model):
    firstname = models.CharField(max_length=100)
    lastname = models.CharField(max_length=100)
    email = models.EmailField(unique=True)
    phone = models.CharField(max_length=32)
    location = models.CharField(max_length=255)
    city = models.CharField(max_length=100)
    country = models.CharField(max_length=100)
    linkedin = models.URLField(blank=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self) -> str:
        return f"{self.firstname} {self.lastname} <{self.email}>"


class Education(models.Model):
    profile = models.ForeignKey(
        UserProfile,
        related_name="educations",
        on_delete=models.CASCADE,
    )
    school = models.CharField(max_length=255)
    field = models.CharField(max_length=255)
    degree = models.CharField(max_length=100)
    start_date = models.DateField()
    end_date = models.DateField()

    def __str__(self) -> str:
        return f"{self.school} - {self.degree}"


class Experience(models.Model):
    profile = models.ForeignKey(
        UserProfile,
        related_name="experiences",
        on_delete=models.CASCADE,
    )
    company = models.CharField(max_length=255)
    title = models.CharField(max_length=255)
    start_date = models.DateField()
    end_date = models.DateField()
    description = models.TextField()

    def __str__(self) -> str:
        return f"{self.company} - {self.title}"

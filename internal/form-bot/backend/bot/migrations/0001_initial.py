import uuid

from django.db import migrations, models


class Migration(migrations.Migration):
    initial = True

    dependencies = []

    operations = [
        migrations.CreateModel(
            name="UserProfile",
            fields=[
                ("id", models.UUIDField(default=uuid.uuid4, editable=False, primary_key=True, serialize=False)),
                ("firstname", models.CharField(max_length=100)),
                ("lastname", models.CharField(max_length=100)),
                ("email", models.EmailField(max_length=254, unique=True)),
                ("phone", models.CharField(blank=True, max_length=32, null=True)),
                ("location", models.CharField(blank=True, max_length=255, null=True)),
                ("city", models.CharField(blank=True, max_length=100, null=True)),
                ("country", models.CharField(blank=True, max_length=100, null=True)),
                ("linkedin", models.URLField(blank=True)),
                ("created_at", models.DateTimeField(auto_now_add=True)),
                ("updated_at", models.DateTimeField(auto_now=True)),
            ],
            options={"db_table": "user_profile"},
        ),
        migrations.CreateModel(
            name="Education",
            fields=[
                ("id", models.UUIDField(default=uuid.uuid4, editable=False, primary_key=True, serialize=False)),
                ("school", models.CharField(max_length=255)),
                ("field", models.CharField(blank=True, max_length=255, null=True)),
                ("degree", models.CharField(blank=True, max_length=100, null=True)),
                ("start_date", models.DateField(blank=True, null=True)),
                ("end_date", models.DateField(blank=True, null=True)),
                ("created_at", models.DateTimeField(auto_now_add=True)),
                (
                    "profile",
                    models.ForeignKey(
                        db_column="user_id",
                        on_delete=models.deletion.CASCADE,
                        related_name="educations",
                        to="bot.userprofile",
                    ),
                ),
            ],
            options={"db_table": "user_education"},
        ),
        migrations.CreateModel(
            name="Experience",
            fields=[
                ("id", models.UUIDField(default=uuid.uuid4, editable=False, primary_key=True, serialize=False)),
                ("company", models.CharField(max_length=255)),
                ("title", models.CharField(blank=True, max_length=255, null=True)),
                ("start_date", models.DateField(blank=True, null=True)),
                ("end_date", models.DateField(blank=True, null=True)),
                ("description", models.TextField(blank=True, null=True)),
                ("created_at", models.DateTimeField(auto_now_add=True)),
                (
                    "profile",
                    models.ForeignKey(
                        db_column="user_id",
                        on_delete=models.deletion.CASCADE,
                        related_name="experiences",
                        to="bot.userprofile",
                    ),
                ),
            ],
            options={"db_table": "user_experience"},
        ),
    ]

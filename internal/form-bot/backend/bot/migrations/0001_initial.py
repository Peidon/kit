from django.db import migrations, models


class Migration(migrations.Migration):
    initial = True

    dependencies = []

    operations = [
        migrations.CreateModel(
            name="UserProfile",
            fields=[
                ("id", models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name="ID")),
                ("firstname", models.CharField(max_length=100)),
                ("lastname", models.CharField(max_length=100)),
                ("email", models.EmailField(max_length=254, unique=True)),
                ("phone", models.CharField(max_length=32)),
                ("location", models.CharField(max_length=255)),
                ("city", models.CharField(max_length=100)),
                ("country", models.CharField(max_length=100)),
                ("linkedin", models.URLField(blank=True)),
                ("created_at", models.DateTimeField(auto_now_add=True)),
                ("updated_at", models.DateTimeField(auto_now=True)),
            ],
        ),
        migrations.CreateModel(
            name="Education",
            fields=[
                ("id", models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name="ID")),
                ("school", models.CharField(max_length=255)),
                ("field", models.CharField(max_length=255)),
                ("degree", models.CharField(max_length=100)),
                ("start_date", models.DateField()),
                ("end_date", models.DateField()),
                (
                    "profile",
                    models.ForeignKey(on_delete=models.deletion.CASCADE, related_name="educations", to="bot.userprofile"),
                ),
            ],
        ),
        migrations.CreateModel(
            name="Experience",
            fields=[
                ("id", models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name="ID")),
                ("company", models.CharField(max_length=255)),
                ("title", models.CharField(max_length=255)),
                ("start_date", models.DateField()),
                ("end_date", models.DateField()),
                ("description", models.TextField()),
                (
                    "profile",
                    models.ForeignKey(on_delete=models.deletion.CASCADE, related_name="experiences", to="bot.userprofile"),
                ),
            ],
        ),
    ]

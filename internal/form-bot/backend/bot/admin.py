from django.contrib import admin

# Register your models here.
from .models import UserProfile, Education, Experience

admin.site.register(UserProfile)
admin.site.register(Education)
admin.site.register(Experience)
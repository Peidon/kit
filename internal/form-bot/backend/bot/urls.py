from django.urls import path

from .views import save_user_info


urlpatterns = [
    path("user_info", save_user_info, name="save user_info"),
]

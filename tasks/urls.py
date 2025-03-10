from django.urls import path
from rest_framework.urlpatterns import format_suffix_patterns
from tasks import views

urlpatterns = [
    path('check-status/', views.CheckStatus.as_view()),
    path('run-soccerway/<int:parser_key>', views.RunSoccerway.as_view())
]

urlpatterns = format_suffix_patterns(urlpatterns)
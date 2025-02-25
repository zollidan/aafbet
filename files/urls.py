from django.urls import path
from rest_framework.urlpatterns import format_suffix_patterns
from files import views

urlpatterns = [
    path('', views.FileList.as_view()),
    path('<str:pk>/', views.FileDetail.as_view()),
    path('s3/download/<str:file_id>/', views.S3FileDownloadView.as_view(), name='s3-file-download'),
    path('s3/list/', views.S3FileListView.as_view(), name='s3-file-list'),
    path('s3/upload/', views.S3FileUploadView.as_view(), name='s3-file-upload')
]

urlpatterns = format_suffix_patterns(urlpatterns)
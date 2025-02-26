import json
import os
import uuid
from django.http import HttpResponse
from rest_framework.response import Response
from rest_framework import generics, status
from aafbet.settings import AWS_BUCKET
from files.models import File
from files.serializers import FileSerializer
from rest_framework.views import APIView
from aafbet.utils import s3_client
from dotenv import load_dotenv
from rest_framework.parsers import MultiPartParser, FileUploadParser

load_dotenv()

class FileList(generics.ListCreateAPIView):
    queryset = File.objects.all()
    serializer_class = FileSerializer
    
class FileDetail(generics.RetrieveUpdateDestroyAPIView):
    queryset = File.objects.all()
    serializer_class = FileSerializer
    
    
class S3FileListView(APIView):
    def get(self, request):
        
        try:
            files_list = s3_client.list_objects(Bucket=AWS_BUCKET)['Contents']

            return Response(files_list, status=status.HTTP_200_OK)
        except Exception as e:
             
            return Response(str(e), status=status.HTTP_500_INTERNAL_SERVER_ERROR)
    
    
class S3FileUploadView(APIView):
    parser_classes = [MultiPartParser]

    def post(self, request):
        try:
            uploaded_file = request.FILES.get('file')
            file_key = f"{uuid.uuid4()}_{uploaded_file.name}"

            if not uploaded_file:
                return Response(
                    {"detail": "No file uploaded"}, 
                    status=status.HTTP_400_BAD_REQUEST
                )
            
            # сделать чтобы имя в s3 было случайным, а в бд сохранялост название и называние в s3
            s3_client.put_object(
                Bucket=AWS_BUCKET,
                Key=file_key,
                Body=uploaded_file.read(),
                ContentType=uploaded_file.content_type
            )
            
            # Сохраняем метаданные в базу данных
            file_data = {
                'name': uploaded_file.name,
                'file_url': "https://aaf-bet.ru/api/files/s3/download/" + file_key,
                "parser_group": "M",
                "s3_key": file_key
            }
            
            serializer = FileSerializer(data=file_data)
            if serializer.is_valid():
                serializer.save()
                return Response(serializer.data, status=status.HTTP_201_CREATED)
            
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

        except Exception as e:
            return Response(
                {"detail": str(e)}, 
                status=status.HTTP_500_INTERNAL_SERVER_ERROR
            )
        

        
    
class S3FileDownloadView(APIView):
    def get(self, request, file_id):
        
        get_object_response = s3_client.get_object(Bucket=AWS_BUCKET, Key=file_id)
        file_binary = get_object_response['Body'].read()
        content_type = get_object_response['ContentType']

        return HttpResponse(file_binary, content_type=content_type)

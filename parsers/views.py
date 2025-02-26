from django.shortcuts import render
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
import requests

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36'
}

class CheckStatus(APIView):
    def get(self, request):
        
        try:
            marafon_response = requests.get('https://www.marathonbet.ru/su/', headers=headers)
            soccerway_response = requests.get('https://www.soccerway.com/', headers=headers)


            return Response({"marafon status": marafon_response.status_code, "soccerway status": soccerway_response.status_code}, status=status.HTTP_200_OK)
        
        except Exception as e:
            return Response(str(e), status=status.HTTP_500_INTERNAL_SERVER_ERROR)

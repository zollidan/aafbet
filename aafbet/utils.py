import boto3
import os
from dotenv import load_dotenv

load_dotenv()

s3_client = boto3.client(
    's3',
    endpoint_url='https://storage.yandexcloud.net',
    aws_access_key_id=os.getenv("AWS_ID"),
    aws_secret_access_key=os.getenv("AWS_SECRET_KEY")
)
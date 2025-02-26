import boto3
import os

from aafbet.settings import AWS_ID, AWS_SECRET_KEY

s3_client = boto3.client(
    's3',
    endpoint_url='https://storage.yandexcloud.net',
    aws_access_key_id=AWS_ID,
    aws_secret_access_key=AWS_SECRET_KEY
)
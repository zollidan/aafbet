import os
from celery import Celery
from parsers.parser1 import run_fake_parser

celery_app = Celery(
    "tasks",
)

celery_app.conf.broker_url = os.environ.get("CELERY_BROKER_URL", "redis://localhost:6379")
celery_app.conf.result_backend = os.environ.get("CELERY_RESULT_BACKEND", "redis://localhost:6379")


celery_app.conf.update(
    task_serializer='json',
    result_serializer='json',
    accept_content=['json'],
)

celery_app.autodiscover_tasks(['app.tasks'])

@celery_app.task
def run_parser_task():
    run_fake_parser()
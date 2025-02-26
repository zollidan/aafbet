import time

from celery import shared_task


@shared_task
def parser1(task_type):

 time.sleep(int(task_type) * 10)

 return True

@shared_task
def parser2(task_type):

 time.sleep(int(task_type) * 10)

 return True
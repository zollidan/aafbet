from django.db import models
import uuid

enumerate

class File(models.Model):
    PARSERS_GROUP = (
    ('1', 'soccerway_1'),
    ('2', 'soccerway_2'),
    ('M', 'marafon')
    )
    
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    name = models.CharField(max_length=100, blank=True, default='')
    file_url = models.URLField(max_length=200, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    parser_group = models.CharField(max_length=1, choices=PARSERS_GROUP)
    
    class Meta:
        ordering = ['created']

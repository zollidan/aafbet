# Generated by Django 5.1.6 on 2025-02-25 18:01

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('files', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='file',
            name='s3_key',
            field=models.CharField(blank=True, default='', max_length=50),
        ),
        migrations.AlterField(
            model_name='file',
            name='parser_group',
            field=models.CharField(choices=[('1', 'soccerway_1'), ('2', 'soccerway_2'), ('M', 'marafon'), ('O', 'other')], max_length=1),
        ),
    ]

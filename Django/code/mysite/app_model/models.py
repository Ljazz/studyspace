from django.db import models


class Artist(models.Model):
    name = models.CharField(max_length=10)


class Album(models.Model):
    artist = models.ForeignKey(
        Artist,
        on_delete=models.CASCADE
    )


class Song(models.Model):
    artist = models.ForeignKey(
        Artist,
        on_delete=models.CASCADE
    )
    album = models.ForeignKey(
        Album,
        on_delete=models.RESTRICT
    )

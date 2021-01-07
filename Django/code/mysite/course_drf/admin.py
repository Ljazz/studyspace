from django.contrib import admin
from course_drf.models import Course

@admin.register(Course)
class CourseAdmin(admin.ModelAdmin):
    list_display = ('name', 'introduction', 'teacher', 'price')
    search_fields = list_display
    list_fitler = list_display

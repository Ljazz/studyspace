from django.shortcuts import render
import time
from ali_pay.models import Pay
from django.views.decorators.csrf import csrf_exempt
from django.http import JsonResponse
from tools.alitool import create_paypage_url, check_pay_state

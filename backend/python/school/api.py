from rest_framework import status
from rest_framework.decorators import api_view
from rest_framework.response import Response


@api_view()
def index(request):
    """
    Returns the alive status of the service.
    """
    return Response(status=status.HTTP_204_NO_CONTENT)


@api_view()
def health(request):
    """
    Returns the alive status of the service.
    """
    return Response({"alive": "ok"})

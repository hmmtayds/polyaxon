# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

from django.conf.urls import url
from rest_framework.urlpatterns import format_suffix_patterns

from libs.urls import UUID_PATTERN
from repos import views

urlpatterns = [
    url(r'^repos/?$',
        views.RepoListView.as_view()),
    url(r'^repos/{}/?$'.format(UUID_PATTERN),
        views.RepoDetailView.as_view()),
    url(r'^repos/{}/upload/?$'.format(UUID_PATTERN),
        views.UploadFilesView.as_view()),
]

urlpatterns = format_suffix_patterns(urlpatterns)

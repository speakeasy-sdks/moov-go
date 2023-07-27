"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import enrichmentaddress as shared_enrichmentaddress
from typing import Optional



@dataclasses.dataclass
class GetEnrichmentAddressRequest:
    search: str = dataclasses.field(metadata={'query_param': { 'field_name': 'search', 'style': 'form', 'explode': True }})
    r"""Partial or complete address to search."""
    exclude_states: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'excludeStates', 'style': 'form', 'explode': True }})
    r"""Exclude list of states from results. No `include` pararmeters may be used with this parameter."""
    include_cities: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'includeCities', 'style': 'form', 'explode': True }})
    r"""Limits results to a list of given cities."""
    include_states: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'includeStates', 'style': 'form', 'explode': True }})
    r"""Limits results to a list of given states."""
    include_zipcodes: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'includeZipcodes', 'style': 'form', 'explode': True }})
    r"""Limits results to a list of given zipcodes."""
    max_results: Optional[int] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'maxResults', 'style': 'form', 'explode': True }})
    r"""Maximum number of results to return."""
    prefer_cities: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'preferCities', 'style': 'form', 'explode': True }})
    r"""Display results with the listed cities at the top."""
    prefer_geolocation: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'preferGeolocation', 'style': 'form', 'explode': True }})
    r"""If omitted or set to `city` it uses the sender's IP address to determine location, then automatically adds the city and state to the preferCities value. This parameter takes precedence over other `include` or `exclude` parameters meaning that if it is not set to `none` you may see addresses from areas you do not wish to see."""
    prefer_ratio: Optional[int] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'preferRatio', 'style': 'form', 'explode': True }})
    r"""Specifies the percentage of address suggestions that should be preferred and will appear at the top of the results."""
    prefer_states: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'preferStates', 'style': 'form', 'explode': True }})
    r"""Display results with the listed states at the top."""
    prefer_zipcodes: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'preferZipcodes', 'style': 'form', 'explode': True }})
    r"""Display results with the listed zipcodes at the top."""
    selected: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'selected', 'style': 'form', 'explode': True }})
    r"""Useful for narrowing results with `addressLine2` suggestions such as `Apt` (denotes an apartment building with multiple residences)."""
    source: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'source', 'style': 'form', 'explode': True }})
    r"""Include results from alternate data sources. Allowed values are -- `all` (non-postal addresses) or `postal` (postal addresses only)."""
    




@dataclasses.dataclass
class GetEnrichmentAddressResponse:
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    enrichment_addresses: Optional[list[shared_enrichmentaddress.EnrichmentAddress]] = dataclasses.field(default=None)
    r"""Address suggestions"""
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    


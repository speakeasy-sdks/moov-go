"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import enrichedbusiness as shared_enrichedbusiness
from dataclasses_json import Undefined, dataclass_json
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class EnrichmentProfile:
    r"""Describes an enriched business profile"""
    business: Optional[shared_enrichedbusiness.EnrichedBusiness] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('business'), 'exclude': lambda f: f is None }})
    r"""Describes a company"""
    


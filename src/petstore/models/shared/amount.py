"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from dataclasses_json import Undefined, dataclass_json
from petstore import utils


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class Amount:
    r"""A representation of money containing an integer value and it's currency."""
    currency: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('currency') }})
    r"""A 3-letter ISO 4217 currency code"""
    value: int = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('value') }})
    r"""Quantity in the smallest unit of the specified currency. In USD this is cents, so $12.04 is 1204 and $0.99 would be 99."""
    


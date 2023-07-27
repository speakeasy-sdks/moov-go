"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import underwritingstatus as shared_underwritingstatus
from dataclasses_json import Undefined, dataclass_json
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class Underwriting:
    r"""Describes underwriting values (in USD) used for card payment acceptance"""
    average_monthly_transaction_volume: Optional[int] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('averageMonthlyTransactionVolume'), 'exclude': lambda f: f is None }})
    average_transaction_size: Optional[int] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('averageTransactionSize'), 'exclude': lambda f: f is None }})
    max_transaction_size: Optional[int] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('maxTransactionSize'), 'exclude': lambda f: f is None }})
    status: Optional[shared_underwritingstatus.UnderwritingStatus] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('status'), 'exclude': lambda f: f is None }})
    r"""The status of underwriting for an account"""
    


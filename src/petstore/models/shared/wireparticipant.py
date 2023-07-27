"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import wirelocation as shared_wirelocation
from dataclasses_json import Undefined, dataclass_json
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class WireParticipant:
    r"""Financial institution information regarding a wire participant"""
    book_entry_securities_transfer_status: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('bookEntrySecuritiesTransferStatus'), 'exclude': lambda f: f is None }})
    customer_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('customerName'), 'exclude': lambda f: f is None }})
    date_: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('date'), 'exclude': lambda f: f is None }})
    funds_settlement_only_status: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('fundsSettlementOnlyStatus'), 'exclude': lambda f: f is None }})
    funds_transfer_status: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('fundsTransferStatus'), 'exclude': lambda f: f is None }})
    location: Optional[shared_wirelocation.WireLocation] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('location'), 'exclude': lambda f: f is None }})
    routing_number: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('routingNumber'), 'exclude': lambda f: f is None }})
    telegraphic_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('telegraphicName'), 'exclude': lambda f: f is None }})
    


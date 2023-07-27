"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import dateutil.parser
from ..shared import amount as shared_amount
from ..shared import failurereason as shared_failurereason
from ..shared import getdispute as shared_getdispute
from ..shared import getfacilitatorfee as shared_getfacilitatorfee
from ..shared import getrefund as shared_getrefund
from ..shared import gettransferfullsource as shared_gettransferfullsource
from ..shared import gettransferfullsourcedestination as shared_gettransferfullsourcedestination
from ..shared import transferstatus as shared_transferstatus
from dataclasses_json import Undefined, dataclass_json
from datetime import datetime
from marshmallow import fields
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class GetTransferFullDisputedAmount:
    r"""A representation of money containing an integer value and it's currency."""
    currency: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('currency') }})
    r"""A 3-letter ISO 4217 currency code"""
    value: int = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('value') }})
    r"""Quantity in the smallest unit of the specified currency. In USD this is cents, so $12.04 is 1204 and $0.99 would be 99."""
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class GetTransferFullRefundedAmount:
    r"""A representation of money containing an integer value and it's currency."""
    currency: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('currency') }})
    r"""A 3-letter ISO 4217 currency code"""
    value: int = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('value') }})
    r"""Quantity in the smallest unit of the specified currency. In USD this is cents, so $12.04 is 1204 and $0.99 would be 99."""
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class GetTransferFull:
    r"""Transfer details"""
    amount: Optional[shared_amount.Amount] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('amount'), 'exclude': lambda f: f is None }})
    r"""A representation of money containing an integer value and it's currency."""
    created_on: Optional[datetime] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('createdOn'), 'encoder': utils.datetimeisoformat(True), 'decoder': dateutil.parser.isoparse, 'mm_field': fields.DateTime(format='iso'), 'exclude': lambda f: f is None }})
    description: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('description'), 'exclude': lambda f: f is None }})
    r"""A description of the transfer"""
    destination: Optional[shared_gettransferfullsourcedestination.GetTransferFullSourceDestination] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('destination'), 'exclude': lambda f: f is None }})
    disputed_amount: Optional[GetTransferFullDisputedAmount] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('disputedAmount'), 'exclude': lambda f: f is None }})
    r"""The total disputed amount for a card transfer"""
    disputes: Optional[list[shared_getdispute.GetDispute]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('disputes'), 'exclude': lambda f: f is None }})
    r"""A list of disputes for a card transfer"""
    facilitator_fee: Optional[shared_getfacilitatorfee.GetFacilitatorFee] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('facilitatorFee'), 'exclude': lambda f: f is None }})
    r"""Fee you charged your customer for the transfer"""
    failure_reason: Optional[shared_failurereason.FailureReason] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('failureReason'), 'exclude': lambda f: f is None }})
    r"""Transfer failure reason"""
    group_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('groupID'), 'exclude': lambda f: f is None }})
    metadata: Optional[dict[str, str]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('metadata'), 'exclude': lambda f: f is None }})
    r"""Free-form key-value pair list. Useful for storing information that is not captured elsewhere."""
    moov_fee: Optional[int] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('moovFee'), 'exclude': lambda f: f is None }})
    r"""Fee charged to your platform account for card transfers"""
    refunded_amount: Optional[GetTransferFullRefundedAmount] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('refundedAmount'), 'exclude': lambda f: f is None }})
    r"""The total refunded amount for a card transfer"""
    refunds: Optional[list[shared_getrefund.GetRefund]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('refunds'), 'exclude': lambda f: f is None }})
    r"""A list of refunds for a card transfer"""
    source: Optional[shared_gettransferfullsource.GetTransferFullSource] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('source'), 'exclude': lambda f: f is None }})
    status: Optional[shared_transferstatus.TransferStatus] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('status'), 'exclude': lambda f: f is None }})
    r"""Current status of a transfer"""
    transfer_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('transferID'), 'exclude': lambda f: f is None }})
    r"""UUID v4"""
    


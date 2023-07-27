"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import accountverificationstatus as shared_accountverificationstatus
from ..shared import document as shared_document
from ..shared import verificationstatus as shared_verificationstatus
from ..shared import verificationstatusdetails as shared_verificationstatusdetails
from dataclasses_json import Undefined, dataclass_json
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class Verification:
    r"""Describes identity verification status and relevant identity verification documents"""
    status: shared_verificationstatus.VerificationStatus = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('status') }})
    r"""This field is deprecated but available for use until February 2023.

    Deprecated: this field will be removed in a future release, please migrate away from it as soon as possible
    """
    details: Optional[shared_verificationstatusdetails.VerificationStatusDetails] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('details'), 'exclude': lambda f: f is None }})
    r"""This field is deprecated but available for use until February 2023.

    Deprecated: this field will be removed in a future release, please migrate away from it as soon as possible
    """
    documents: Optional[list[shared_document.Document]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('documents'), 'exclude': lambda f: f is None }})
    verification_status: Optional[shared_accountverificationstatus.AccountVerificationStatus] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('verificationStatus'), 'exclude': lambda f: f is None }})
    r"""The status of an identity verification for a profile"""
    


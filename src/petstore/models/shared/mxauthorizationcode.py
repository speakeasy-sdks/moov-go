"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from dataclasses_json import Undefined, dataclass_json
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class MXAuthorizationCode:
    r"""The authorization code of a MX account which allows a processor to retrieve a linked payment account. <br><br> `sandbox` - When linking a bank account to a `sandbox` account using a MX authorization code it will utilize MX's sandbox environment. The MX authorization code provided must be generated from MX's sandbox environment."""
    authorization_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('authorizationCode'), 'exclude': lambda f: f is None }})
    


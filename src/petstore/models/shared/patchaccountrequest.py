"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import ein as shared_ein
from dataclasses_json import Undefined, dataclass_json
from enum import Enum
from petstore import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestCustomerSupportAddress:
    address_line1: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine1'), 'exclude': lambda f: f is None }})
    address_line2: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine2'), 'exclude': lambda f: f is None }})
    city: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('city'), 'exclude': lambda f: f is None }})
    country: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('country'), 'exclude': lambda f: f is None }})
    postal_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('postalCode'), 'exclude': lambda f: f is None }})
    state_or_province: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('stateOrProvince'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestCustomerSupportPhone:
    country_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('countryCode'), 'exclude': lambda f: f is None }})
    number: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('number'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestCustomerSupport:
    r"""User-provided information that can be displayed on credit card transactions for customers to use when contacting a customer support team. This data is only allowed on a business account"""
    address: Optional[PatchAccountRequestCustomerSupportAddress] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('address'), 'exclude': lambda f: f is None }})
    email: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('email'), 'exclude': lambda f: f is None }})
    r"""Email Address"""
    phone: Optional[PatchAccountRequestCustomerSupportPhone] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('phone'), 'exclude': lambda f: f is None }})
    website: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('website'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileBusinessAddress:
    address_line1: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine1'), 'exclude': lambda f: f is None }})
    address_line2: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine2'), 'exclude': lambda f: f is None }})
    city: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('city'), 'exclude': lambda f: f is None }})
    country: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('country'), 'exclude': lambda f: f is None }})
    postal_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('postalCode'), 'exclude': lambda f: f is None }})
    state_or_province: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('stateOrProvince'), 'exclude': lambda f: f is None }})
    


class PatchAccountRequestProfileBusinessBusinessType(str, Enum):
    r"""The type of entity represented by this Business"""
    SOLE_PROPRIETORSHIP = 'soleProprietorship'
    UNINCORPORATED_ASSOCIATION = 'unincorporatedAssociation'
    TRUST = 'trust'
    PUBLIC_CORPORATION = 'publicCorporation'
    PRIVATE_CORPORATION = 'privateCorporation'
    LLC = 'llc'
    PARTNERSHIP = 'partnership'
    UNINCORPORATED_NON_PROFIT = 'unincorporatedNonProfit'
    INCORPORATED_NON_PROFIT = 'incorporatedNonProfit'


@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileBusinessIndustryCodes:
    r"""Describes industry specific identifiers"""
    mcc: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('mcc'), 'exclude': lambda f: f is None }})
    naics: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('naics'), 'exclude': lambda f: f is None }})
    sic: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('sic'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileBusinessPhone:
    country_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('countryCode'), 'exclude': lambda f: f is None }})
    number: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('number'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileBusinessTaxID:
    r"""An EIN (employer identification number) for the business. For sole proprietors, an SSN can be used as the EIN."""
    ein: Optional[shared_ein.Ein] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('ein'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileBusiness:
    r"""Describes the fields available when patching a business"""
    address: Optional[PatchAccountRequestProfileBusinessAddress] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('address'), 'exclude': lambda f: f is None }})
    business_type: Optional[PatchAccountRequestProfileBusinessBusinessType] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('businessType'), 'exclude': lambda f: f is None }})
    description: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('description'), 'exclude': lambda f: f is None }})
    doing_business_as: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('doingBusinessAs'), 'exclude': lambda f: f is None }})
    email: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('email'), 'exclude': lambda f: f is None }})
    industry_codes: Optional[PatchAccountRequestProfileBusinessIndustryCodes] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('industryCodes'), 'exclude': lambda f: f is None }})
    legal_business_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('legalBusinessName'), 'exclude': lambda f: f is None }})
    owners_provided: Optional[bool] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('ownersProvided'), 'exclude': lambda f: f is None }})
    phone: Optional[PatchAccountRequestProfileBusinessPhone] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('phone'), 'exclude': lambda f: f is None }})
    tax_id: Optional[PatchAccountRequestProfileBusinessTaxID] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('taxID'), 'exclude': lambda f: f is None }})
    website: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('website'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualAddress:
    address_line1: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine1'), 'exclude': lambda f: f is None }})
    address_line2: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('addressLine2'), 'exclude': lambda f: f is None }})
    city: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('city'), 'exclude': lambda f: f is None }})
    country: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('country'), 'exclude': lambda f: f is None }})
    postal_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('postalCode'), 'exclude': lambda f: f is None }})
    state_or_province: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('stateOrProvince'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualBirthDate:
    r"""Birthdate for an individual"""
    day: int = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('day') }})
    month: int = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('month') }})
    year: int = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('year') }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualGovernmentIDItin:
    full: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('full'), 'exclude': lambda f: f is None }})
    last_four: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('lastFour'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualGovernmentIDSsn:
    full: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('full'), 'exclude': lambda f: f is None }})
    last_four: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('lastFour'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualGovernmentID:
    itin: Optional[PatchAccountRequestProfileIndividualGovernmentIDItin] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('itin'), 'exclude': lambda f: f is None }})
    ssn: Optional[PatchAccountRequestProfileIndividualGovernmentIDSsn] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('ssn'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualName:
    r"""Name for an individual"""
    first_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('firstName'), 'exclude': lambda f: f is None }})
    r"""Name this person was given. This is usually the the same as first name."""
    last_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('lastName'), 'exclude': lambda f: f is None }})
    r"""Family name of this person. This is usually the the same as last name."""
    middle_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('middleName'), 'exclude': lambda f: f is None }})
    r"""Name this person was given. This is usually the the same as first name."""
    suffix: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('suffix'), 'exclude': lambda f: f is None }})
    r"""Suffix of a given name"""
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividualPhone:
    country_code: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('countryCode'), 'exclude': lambda f: f is None }})
    number: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('number'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfileIndividual:
    r"""Describes the fields available when patching an individual"""
    address: Optional[PatchAccountRequestProfileIndividualAddress] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('address'), 'exclude': lambda f: f is None }})
    birth_date: Optional[PatchAccountRequestProfileIndividualBirthDate] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('birthDate'), 'exclude': lambda f: f is None }})
    email: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('email'), 'exclude': lambda f: f is None }})
    government_id: Optional[PatchAccountRequestProfileIndividualGovernmentID] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('governmentID'), 'exclude': lambda f: f is None }})
    name: Optional[PatchAccountRequestProfileIndividualName] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('name'), 'exclude': lambda f: f is None }})
    phone: Optional[PatchAccountRequestProfileIndividualPhone] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('phone'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestProfile:
    r"""Describes the fields available when patching a profile.
    Each object can be patched independent of patching the other fields.
    """
    business: Optional[PatchAccountRequestProfileBusiness] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('business'), 'exclude': lambda f: f is None }})
    individual: Optional[PatchAccountRequestProfileIndividual] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('individual'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestSettingsAchPayment:
    r"""User provided settings to manage ACH payments"""
    company_name: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('companyName'), 'exclude': lambda f: f is None }})
    r"""The description that shows up on ACH transactions. This will default to the account's display name on account creation."""
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestSettingsCardPayment:
    r"""User provided settings to manage card payments. This data is only allowed on a business account"""
    statement_descriptor: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('statementDescriptor'), 'exclude': lambda f: f is None }})
    r"""The description that shows up on credit card transactions. This will default to the accounts display name on account creation."""
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestSettings:
    r"""User provided settings to manage an account"""
    ach_payment: Optional[PatchAccountRequestSettingsAchPayment] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('achPayment'), 'exclude': lambda f: f is None }})
    card_payment: Optional[PatchAccountRequestSettingsCardPayment] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('cardPayment'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequestTermsOfService:
    r"""An encrypted value used to record acceptance of Moov's Terms of Service"""
    token: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('token'), 'exclude': lambda f: f is None }})
    



@dataclass_json(undefined=Undefined.EXCLUDE)

@dataclasses.dataclass
class PatchAccountRequest:
    r"""Describes the fields available when patching a Moov Account"""
    customer_support: Optional[PatchAccountRequestCustomerSupport] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('customerSupport'), 'exclude': lambda f: f is None }})
    foreign_id: Optional[str] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('foreignID'), 'exclude': lambda f: f is None }})
    metadata: Optional[dict[str, str]] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('metadata'), 'exclude': lambda f: f is None }})
    r"""Free-form key-value pair list. Useful for storing information that is not captured elsewhere."""
    profile: Optional[PatchAccountRequestProfile] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('profile'), 'exclude': lambda f: f is None }})
    settings: Optional[PatchAccountRequestSettings] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('settings'), 'exclude': lambda f: f is None }})
    terms_of_service: Optional[PatchAccountRequestTermsOfService] = dataclasses.field(default=None, metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('termsOfService'), 'exclude': lambda f: f is None }})
    


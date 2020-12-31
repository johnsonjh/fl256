# fl256

## A high-precision fixed-point numeric representation library for Go

[![License-FL256](https://img.shields.io/badge/License-fl256-blue.svg)](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.256)
[![License-GFODLv1](https://img.shields.io/badge/License-OX153MkI-red.svg)](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.153
[![CodebeatBadge](https://codebeat.co/badges/85330b22-1d38-4937-9fd9-c506098f210e)](https://codebeat.co/projects/github-com-johnsonjh-fl256-master)
[![Maintainability](https://api.codeclimate.com/v1/badges/3c7d3ad4cb992d2cd80d/maintainability)](https://codeclimate.com/github/johnsonjh/fl256/maintainability)

-----------------------

### Overview

`fl256` implements a high-precision fixed-point numeric type for Go.
It may be useful in certain domain-specific mathematical contexts.

### Implementation

**NOTE**: This library is built on **big.Float**. As such, it does **not**
implement *end-to-end* fixed-point calculations, and is **not** using BCD
or other native decimal types. It was originally intended for use in a
crypto-currency ledger system, with specific constraints: a currency with
a maximum supply of 4,398,046,511,103, represented with an unsigned
fixed-point 42.214 256-bit interchange format, and is not intended for
general purpose decial calculations or accounting applications.

### Alternatives

Potential users *must* fully understand the implications and limitations of
this implementation, and evaluate if this system may be appropriate for a
particular application or use case.

There are many mature alternatives that exist for performing financial
calculations using fixed-point operations for general use, such as:
 - ["shopspring/decimal"](https://github.com/shopspring/decimal), 
 - ["cockroachdb/apd"](https://github.com/cockroachdb/apd), 
 - ["robaho/fixed"](https://github.com/robaho/fixed), 
 - ["maurodelazeri/go-number"](https://github.com/maurodelazeri/go-number), 
 - ["lovung/decimal"](https://github.com/lovung/decimal), 
 - ["ericlagergren/decimal"](https://github.com/ericlagergren/decimal), 
 - ["umran/money"](https://github.com/umran/money), 
 - ["db47h/decimal"](https://github.com/db47h/decimal), 
 - ["wayn3h0/go-decimal"](https://github.com/wayn3h0/go-decimal), 
 - ["strongo/decimal"](https://github.com/strongo/decimal), 
 - ["go-info/inf"](https://github.com/go-inf/inf), 

Please follow standard best practices and exercise due diligence before
selecting and adopting any decimal or arbitrary-precision math library.

## Original Author

* [l0k18 (Loki)](https://github.com/l0k18) [\<stalker.loki@protonmail.ch\>](mailto:stalker.loki@protonmail.ch)

  * The original author released the `float256` library, from which `fl256` was
    derived, into the public domain, by way of the
	[Unlicense](https://unlicense.org)
	declaration and dedicaton, which 
	[allowed for](https://ar.to/2010/01/dissecting-the-unlicense)
	the `fl256` package to be created and relicensed under unambigious terms.

## Licensing

`fl256` is distributed as uniquely *multi-licensed software*.

This package is *dual-licensed*, and made available under both a 
traditional open-source license, as well as a restricted-use, non-free, 
proprietary license.

 * [**The fl256 License**](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.256),
 is the fl256 open-source software license. It is *fully compatible* with
 [**The X.Org Preferred License**](https://gitlab.freedesktop.org/xorg/doc/xorg-docs/-/blob/master/general/License.xml),
 (itself being a minor variant of [**The MIT License**](https://tldrlegal.com/license/mit-license).)
 **The fl256 License** verbiage is, in fact, identical to that of 
 **The X.Org Preferred License**, differentiated only by changes in punctuation.

 * [**The Oxford 153 Entitlement: Mark I**](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.153)
 is a new, unique, non-free, and strongly proprietary software license. While
 the author _strongly_ encourages any eligible user to accept the Terms and
 Conditions of the **OX153MkI**, it is by no means a requirement; anyone who
 will not, or can not, be bound by the terms set forth in the **OX153MkI** 
 may fork this repository and remove all references to the **OX153MkI**.

## Contributing

Anyone who wishes to contribute code or documentation to `fl256` will be
required to digitally sign, using [PGP](https://www.openpgp.org/) or 
[opmsg](https://github.com/stealth/opmsg), the *fl256 CLA/CTA/LO*, a combined
contributor license agreement, copyright transfer assignment, and loyalty oath.
Because of the restrictive terms and conditions of 
**The Oxford 153 Entitlement: Mark I**, every contributor, no matter how small
their contribution, must sign the *fl256 CLA/CTA/LO* document. 

Our Contributor License Agreement, Copyright Transfer Assignment, and Loyalty
Oath is intended to ensure, forever, the integrity of any tool distributed
under the **OX153MkI**. As such, it is a non-negotiable eternally binding legal
instrument, by which the contributor grants the "*The fl256 Authors*" a
perpetual, world- and universe-wide, non-exclusive, free-of-charge, 
zero-royalty, forever irrevocable, persistent, irrevocable, absolute, 
unalterable, and completely unlimited license, as well as a complete and total
reassignment of copyright and ownership, allowing *The fl256 Authors* to
reproduce, redistribute, sub-license, re-license, or otherwise fully control
the contribution, in it's entirety, including any works which may be later
derived from the contribution. The document also includes a Loyalty Oath,
intended to formally and authoritatively finalize the total acceptance of,
and assure eternal adherence to, the principles and values embodied by the
**The Oxford 153 Entitlement: Mark I**.

A complete copy of *fl256 CLA/CTA/LO* will be furnished to any potential
contributor. We encourage that all potential contributors completely
understand the *fl256 CLA/CTA/LO*, and recommended seeking guidance from
qualified legal council well-versed in such matters, who can offer any
necessary assistance and clarification if any questions arise during the
reading of the *fl256 CLA/CTA/LO* document.

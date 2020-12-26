# fl256

## A high-precision fixed-point numeric representation library for Go

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
    derived, into the public domain, by way of the [Unlicense](https://unlicense.org)
	declaration and dedicaton, which [allowed for](https://ar.to/2010/01/dissecting-the-unlicense)
	the `fl256` package to be relicensed under unambigious terms.

## Licensing

`fl256` is distributed as uniquely *multi-licensed software*.

This package is *dual-licensed*, and made available under both a traditional
open-source license, as well as a restricted-use, non-free, proprietary license.

 * **The fl256 License**, is the open-source software license. It is *fully compatible*
with [**The X.Org Preferred License**](https://gitlab.freedesktop.org/xorg/doc/xorg-docs/-/blob/master/general/License.xml),
(itself being a minor variant of [**The MIT License**](https://tldrlegal.com/license/mit-license).)
**The fl256 License** verbiage is, in fact, identical to that of 
**The X.Org Preferred License**, differentiated only by changes in punctuation.

 * [**The Oxford 153 Entitlement: Mark I**](https://prone.ws/ox153) is a new, unique,
non-free, and proprietary software license. While the author _strongly_ encourages all
eligible users to accept the Terms and Conditions of the **OX153MkI**, it is by no
means a requirement. Anyone who will not, or can not, be bound by the terms set forth
in the **OX153MkI** may fork this repository and remove all references to the **OX153MkI**.

## Contributing

Anyone who wishes to contribute code or documentation to `fl256` will be required to
digitally sign, using PGP, Opmsg, or S/MIME, the fl256 *CLA/CTA*, a contributor license
agreement and copyright transfer assignment. Because of the restrictive terms and
conditions of **The Oxford 153 Entitlement: Mark I**, every contributor, no matter
how small their contribution, must sign the fl256 CLA/CTA. 

Our Contributor License Agreement and Copyright Transfer Assignment is intended to
ensure the integrity of any distributed under the **OX153MkI**. As such, it is a
non-negotiable and eternally binding legal instrument, which the contributor uses to
grant the "*The fl256 Authors*" a perpetual, world- and universe-wide, non-exclusive,
free of charge, zero-royalty, forever irrevocable, and completely unlimited license,
as well as a complete and total reassignment of copyright and ownership, allowing us
to reproduce, redistribute, sub-license, re-license, or otherwise control, in its'
entirety, any all facets of the contribution provided, including any works which may
be later derived from the original contribution. The document also includes a loyalty
oath and a renunciation and abdication of the personal sovereignty of the contributor,
intended to assure eternal adherence and formally codify the acceptance of the principles
and values embodied by the **The Oxford 153 Entitlement: Mark I**.

A complete copy of our fl256 CLA/CTA will be furnished to any potential contributors.
We encourage that you completely understand the document, and urge you to seek out 
qualified legal council well-versed in such matters, to offer any necessary assistance
or answer any questions which may arise from the your reading the fl256 CLA/CTA documents.

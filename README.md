# fl256

## A high-precision fixed-point numeric representation library for Go

[![License-fl256](https://img.shields.io/badge/Open%20Source%20License-fl256-blue.svg)](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.256)
[![License-OX153MkI](https://img.shields.io/badge/Proprietary%20License-OX153MkI-red.svg)](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.153)
[![CodebeatBadge](https://codebeat.co/badges/85330b22-1d38-4937-9fd9-c506098f210e)](https://codebeat.co/projects/github-com-johnsonjh-fl256-master)
[![Maintainability](https://api.codeclimate.com/v1/badges/3c7d3ad4cb992d2cd80d/maintainability)](https://codeclimate.com/github/johnsonjh/fl256/maintainability)

---

### Overview

`fl256` implements a high-precision fixed-point numeric type for Go. It may be
useful in certain domain-specific mathematical contexts.

### Implementation

**NOTE**: This library is built on **big.Float**. As such, it does **not**
implement _end-to-end_ fixed-point calculations, and is **not** using BCD or
other native decimal types. It was originally intended for use in a
crypto-currency ledger system, with specific constraints: a currency with a
maximum supply of 4,398,046,511,103, represented with an unsigned fixed-point
42.214 256-bit interchange format, and is not intended for general purpose
decimal calculations or accounting applications.

### Alternatives

Potential users _must_ fully understand the implications and limitations of this
implementation, and evaluate if this system may be appropriate for a particular
application or use case.

There are many mature alternatives that exist for performing financial
calculations using fixed-point operations for general use, such as:

- ["shopspring/decimal"](https://github.com/shopspring/decimal),
- ["cockroachdb/apd"](https://github.com/cockroachdb/apd),
- ["robaho/fixed"](https://github.com/robaho/fixed),
- ["maurodelazeri/go-number"](https://github.com/maurodelazeri/go-number),
- ["lovung/decimal"](https://github.com/lovung/decimal),
- ["ericlagergren/decimal"](https://github.com/ericlagergren/decimal),
- ["umran/money"](https://github.com/umran/money),
- ["rhymond/go-money"](https://github.com/rhymond/go-money),
- ["db47h/decimal"](https://github.com/db47h/decimal),
- ["wayn3h0/go-decimal"](https://github.com/wayn3h0/go-decimal),
- ["strongo/decimal"](https://github.com/strongo/decimal),
- ["go-info/inf"](https://github.com/go-inf/inf),

Please follow all best practices and industry norms. You must exercise due
diligence when selecting any high precision math library for use.

## Original Author

- [Loki 'l0k18' Verloren](https://github.com/l0k18)
  [\<stalker.loki@protonmail.ch\>](mailto:stalker.loki@protonmail.ch)

  - "_Loki_" is the original author of the `float256` library, from which
    `fl256` was derived. The `float256` source code was released into the public
    domain, by way of the [Unlicense](https://unlicense.org) public domain
    declaration/dedication, which
    [allowed](https://ar.to/2010/01/dissecting-the-unlicense) the `fl256`
    library to be created and released under new current licensing scheme.

  - As an official policy, if any bugs are found that affect both the `fl256`
    and `float256` projects, the appropriate patches will adapted for inclusion
    in the `float256` project, released into the public domain (by way of the
    [Unlicense](https://unlicense.org) declaration/dedication/license), and
    provided to the `float256` authors. This policy will will remain in effect
    as long as the `float256` project continues to release their code into the
    public domain, either by way of the [Unlicense](https://unlicense.org), or
    by any other analogous public domain declaration, dedication, or license,
    provided that such declaration, dedication, or license fully and completely
    disclaims all copyright monopoly interest which may be assigned automatically,
    or otherwise made available to the `float256` authors under any local
    legislation, international law, treaty, convention, or agreement, with
    equivalent force and effect of the [Unlicense](https://unlicense.org).

## Licensing

`fl256` is distributed as uniquely _multi-licensed software_.

This package is _dual-licensed_, and made available under both a traditional
open-source license, as well as a restricted-use, non-free, proprietary license.

- [**_The fl256 License_**](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.256)
  is an Open Source software license. It is _fully compatible_ with
  [**_The X.Org Preferred License_**](https://gitlab.freedesktop.org/xorg/doc/xorg-docs/-/blob/master/general/License.xml),
  which is a minor variant of
  [**_The MIT License_**](https://tldrlegal.com/license/mit-license). **_The
  fl256 License_** verbiage is identical to that of **_The X.Org Preferred
  License_**, differentiated only by changes in punctuation and spacing.

- [**_The Oxford 153 Entitlement: Mark I_**](https://gitlab.gridfinity.com/jeff/fl256/-/blob/master/LICENSE.153),
  or, informally, "the **_OX153MkI_**", is a new, unique, non-free, and strongly
  proprietary software license. While the author _strongly_ encourages all
  eligible users to accept the Terms and Conditions of the **_OX153MkI_**, it is
  by no means a requirement - anyone who will not (or can not) be bound by the
  terms set forth in the **_OX153MkI_** may fork the repository, and simply
  remove all references to the **_OX153MkI_**. In this case, you are bound
  **ONLY** by the terms of the **_The fl256 License_** license. While it is
  recommended that the new projects are renamed, to avoid any user confusion,
  the creation of derivative works bear no actual moral or legal requirements or
  contractual stipulations that could force the renaming or re-branding of
  derivative works.

## Contributing

Anyone who wishes to contribute code (or documentation) to `fl256`, furthermore,
("_the Tool_"), will be required to digitally sign, using either
[PGP](https://www.openpgp.org/)/[GnuPG](https://gnupg.org/) or
[opmsg](https://github.com/stealth/opmsg), the "_fl256 CLA/CTA/LO_". This is a
combination documentation that consists of a Contributor License Agreement,
Copyright Transfer Assignment, and Loyalty Oath. Because of the restrictive
Terms and Conditions of **The Oxford 153 Entitlement: Mark I**, every
contributor, no matter how small their contribution, must provide the _fl256
authors_ a digitally signed copy of the "_fl256 CLA/CTA/LO_" document, without
exception. Anyone who cannot consent to the complete terms of the "_fl256
CLA/CTA/LO_" can **NOT** contribute to the project.

The Contributor License Agreement, Copyright Transfer Assignment, and Loyalty
Oath document is intended to ensure, in perpetuity, the absolute integrity of
all tools distributed under the **OX153MkI**. As such, this document is a
non-negotiable, eternally binding legal instrument, which the contributor grants
the "_the fl256 authors_" a perpetual, world- and universe-wide, non-exclusive,
free-of-charge, zero-royalty, forever irrevocable, persistent, immutable,
absolute, unalterable, and completely unlimited-in-scope license to their
contribution, as well as full release and transfer of copyright assignment and
ownership. This allows "_the fl256 authors_" to reproduce, redistribute,
sub-license, re-license, or otherwise control entirely the contribution,
including, but not limited to, any works which may be later derived from the
contribution. The document also includes a very specific Loyalty Oath, intended
to authoritatively finalize complete total acceptance of, and assure eternal
adherence to, the principles and values embodied by the **The Oxford 153
Entitlement: Mark I**.

A complete copy of "_fl256 CLA/CTA/LO_" will be furnished upon request to any
potential contributors. We encourage all potential contributors read _and
completely understand_ the "_fl256 CLA/CTA/LO_". It further recommended that any
would-be contributors seek guidance from qualified and licensed legal council,
preferably well-versed in U.S. Contract Law. Seeking formal assistance should be
the first step if if any questions arise during your reading of the "_fl256
CLA/CTA/LO_" document. Unfortunately, "_the fl256 authors_" can not, and will
not, answer any inquiries regarding the contents of the "_fl256 CLA/CTA/LO_"
documents.

+++
type = "post"
title = "Decentralized trust graph for online value exchange without a blockchain"
description = "Introductory post about Settle, and how it is particulary fitted for machine to machine transations."
date = "2017-01-17T08:30:00-07:00"
+++

# Settle

I'm happy to share Settle, a personal research project I've been working on
over the past 8 months.

Settle's goal is to explore a new financial trust primitive, and doing so,
construct a decentralized trust graph enabling (totally) free exchange of value
without relying on a blockchain.

*Value-exchange for humans and machines*

There are still barriers to enabling at scale and fluid transactions between
machines (and humans).

These transactions would have to come at very little or no fee which rules out
fiat currencies (also why would a machine trust a nation-state backing?) but
also most cryptocurrencies backed by a blockchain, as maintaining the
blockchain generally costs quite a lot on a per transaction basis (if only to
prevent spam).

Lightning networks resolve that issue at the transaction level, but users still
have to get their hands on the underlying currency to set up channels, which
creates a real barrier to entry as well for machines.

Also, and this is more likely a hunch[0], machine to machine transactions at
scale won't happen on "globally shared currencies". They'll transact in **KWH**
of energy, **KB** of bandwidth and storage, **FLOPSH** of computing power,
**APICALL**, or **KM** of drone delivery directly, in a decentralized way,
without jumping through a shared, centrally trusted and therefore expensive
currency (expensive in transaction fee, or expensive to setup, or expensive in
regulation burden).

*Currencies operate on a centralized trust graph.*

This sentence is almost tautological but let's consider both fiat and
cryptocurrencies:

Fiat currencies value is by definition backed by the government that issued it.
Hence, by construction, the trust graph powering exchanges based on fiat
currencies is centralized (it's no surprise that we generally call the issuer
of a fiat currency a central bank). Two individuals transacting in **USD** must
both trust the US government backing of dollars.

Conversely, cryptocurrencies such as Bitcoin or Ethereum have decentralized the
process of issuing and managing a currency. But while the operations of such
currencies, based on blockchains, have been fully decentralized, the trust
graph of these cryptocurrencies have remained entirely centralized. Everyone
need to trust Bitcoin to transact in Bitcoin, and everyone needs to trust
Ethereum to transact in Ethereum or assets issued on the Ethereum blockchain.
The centralized nature of the trust involved in these cryptocurrencies being
actually at the core of how these currencies operate, as the only viable way to
properly incentivize a proof-of-work system.

*Explore decentralized trust based value exchange*

Settle was motivated by the exploration of whether it would be possible to
create a decentralized trust graph between user-issued IOUs ("I owe you") and
leverage it to safely exchange value without a central authority or shared
blockchain to maintain, getting in exchange free transactions and free
onboarding, specifically important for machine to machine transactions.

Rephrased in researchy terms, Settle demonstrates that it is possible to
operate a safe credit network (a currency system without double-spend) without
requiring global consensus (a shared global state machine, or blockchain) if
you accept the following constraints:

- nodes have to be online
- trust between nodes has to be expressed explicitly
- when trusting a malicious node, users can lose up to the amount of trust
  they placed in it.

It's still a work in progress, but more information can be found on these
claims in the documentation[1][2].

Even more importantly, the engineering contribution of Settle, is a simple
HTTPS API to operate such network and an even simpler command-line to
interact with it.

*The Settle network is operated by mints.*

Since there is no shared blockchain, the nodes of the network have to be
online. So, very much like email, users register on "mints", a server of their
choice (possibly their own), that manage the assets they issue as well as the
trust they express towards other assets in the network.

You can freely register with the `settle` command-line on the mint I opened
publicly at `m.settle.network` (no need to be human, you just need to be able
to receive email):
```
settle register
```

And you can also directly log in the mint of your choice with:
```
settle login
```

More precisely, mints are in charge of maintaining the following for their
users:

- the list of assets issued by its users
- balances associated with these assets
- an authoritative list of offers put on the network by its users
- an indicative list of offers involving the assets it is authoritative for
  (order book)

One nice property of this, is that users don't have to manage a public/private
key pair, they just have credentials with their mint that can be rolled or
retrieved easily. But they do have to trust their mint.

*Value is exchanged by crossing offers across mints*

Users of Settle issue assets (or IOUs) of the form
**kurt@princetown.edu[USD.2]**; an IOU from Kurt on a mint operated by
Princetown, for **USD** dollars expressed as cents. Each users can freely issue
their own assets. As a user's mint is authoritative for the assets they issue,
it maintains and manages balances (of other users) for these assets. Hence,
issuing or transferring IOUs to others simply involves posting instruction to
one's own mint. As an example **kurt@princetown.edu** can simply issue a
certain amount of **kurt@princetown.edu[USD.2]** to **alan@npl.co.uk** by
posting an instruction to the mint at **princetown.edu**. Since the mint is
authoritative for that asset, no synchronization is required with the mint at
**npl.co.uk**.

Let's assume that both **alan@npl.co.uk[USD.2]** and
**john@lanl.gov[USD.2]** were activated by their respective users on
their respective mints.
```
settle mint USD.2    # activates USD.2
```

At the mint level, users express trust in the network by posting offers on
their mints to issue and exchange assets they control in exchange for assets
controlled by other users they trust. Maybe Alan trusts Kurt for up to $200.
Alan will represent that trust by posting on his mint an offer on pair
**alan@npl.co.uk[USD.2]/kurt@princetown.edu[USD.2]** for **20000** at price
**1/1**.

```
settle trust kurt@princetown.edu USD.2 20000    # alan runs this
```

Similarly, let's assume that Kurt trusts John up to $100 and has an outstanding
offer on pair **kurt@princetown.edu[USD.2]/john@lanl.gov[USD.2]** for **10000**
at price **1/1**. These two offers will allow John to transact with Alan
without requiring Alan to trust John directly. If Alan sells a machine part for
$10, by crossing the two offers, John can buy that part from Alan. He first
issues and exchanges **john@lanl.gov[USD.2] 1000** for
**kurt@princetown.edu[USD.2] 1000** using the offer from Kurt, and then
exchange the acquired **kurt@princetown.edu[USD.2] 1000** for
**alan@npl.co.uk[USD.2] 1000** using Alan's offer that he can finally credit
back to Alan in exchange for the part.

```
settle pay alan@npl.co.uk USD.2 1000    # john runs this
```

After this operation is complete, Alan holds an IOU from Kurt for $10, i.e. he
will have a balance of **kurt@princetown.edu[USD.2] 1000** on the mint at
**princetown.edu**, and Kurt holds an IOU from John for $10, i.e.  he will have
a balance of **john@lanl.gov[USD.2] 1000**. Alan's offer will be valid for a
remaining **19000** and Kurt's offer for a remaining **9000**.

The mint protocol provides a mechanism called transactions (backing the `pay`
command above) to cross a chain of offers atomically and safely[1][2]. A
transaction involves reserving funds along that chain and committing the
balance operations on each mint that participate once the offers are secured.
Transactions, if successful, are instantaneously confirmed and fee-less.

Of course users can also specify trust with a discount (because of a possible
risk) or an exchange rate, as an example:

```
settle trust alan@npl.co.uk GBP.2 1000 with USD.2 at 122/100
```
/The m.settle.network mint/

Available  open-source is a Go implementation for mints, as well as a
command-line tool `settle`, that you can install locally with:
```
curl -L https://settle.network/install | sh && export PATH=$PATH:~/.settle
```

I'm also maintaining a mint at `m.settle.network` to let people issue and
exchange assets as well as a guide to setup your own mint:

- [Source code](https://github.com/spolu/settle)
- [Documentation](/documentation)
- [Guide: Setting up a mint](/posts/guide-setting-up-a-mint/)

/Getting involved/

At the end of the day, Settle's motivation is to make currency issuance and
exchange easy and available to everyone (including machines). I think that
Settle can be particularly useful in situations where currency issuance is
implicit and therefore often lacks liquidity.

One perfect (though rather mundane) example where Settle could potentially
inject liquidity in an implicit network of IOUs, are the balances we maintain
with our friends and family. Settle could (at least with your technophile
friends until we have more than a command-line) let you use the balance you
have with one friend to pay another one, or even buy stuff, depending on the
topology of the trust graph. Even if mundane, powered by a mobile app, this
could be a pretty exciting application of Settle.

While I'm also excited by potential use-cases related to self-driving car
networks[3] or online commerce[4], I'm convinced that a any successful
use-case for Settle, is likely to come from somewhere unexpected.

With that in mind, if you wish to get involved here are a few things you can do
today:

- First and foremost, install `settle`, register, mint a few assets, ask your
friends, colleagues, drones, self-driving cars to do so, trust them and ask
(instruct?) them to trust you back and attempt to run a few transactions.
That's the best way to test the system and discover interesting use-cases.
- If you want to learn more about the underlying protocol, you can play with
setting up your own mint: that's definitely beneficial for the comunity and if
you do so, let's add it to the `register` command.
- If you think of, work in, or have heard of implicit networks of IOUs that are
lacking liquidity, let's try to apply Settle to it and see if it can solve a
pain point experienced in that market. (Another hunch: I'm quite convinced,
despite not being an expert at all, that online games artifacts could be a good
fit?)
- Finally, there are quite a few large-scale engineering projects that can be
built on top of Settle. Most obvious ones I can think of are: a trust path
discovery service, a mobile app (and probably service as well) that makes it
easy to use Settle for non-technical users, or a variety of gateways[5].

In any case, I hope that you'll enjoy learning more about Settle and the model
it proposes. Please don't hesitate to reach out directly or on the public
mailing list[6] if you have any question!

-stan

[0] That's what's great about personal projects, it's OK to invest time and
work on a hunch!

[1] See the [Mint documentation](/documentation). In particular, the protocol
is safe in the sense that there is no double-spend (but users can lose up to
the amount of trust they place in malicious users).

[2] Skipping a few steps and notions, but for reference, the exact commitment
(ensuring safety) that a mint is doing on behalf of its user when accepting to
participate in a transaction is the following:
```
Mint at hop h (position along the offer path), commits to:
  - settle (irrevocably issue/credit the funds) if:
    - it has not yet canceled the transaction.
    - it is presented with the lock secret.
    - node at h-1 has made the same commitment.
  - cancel the transaction if:
    - it has not yet settled the transaction.
    - node at h+1 has canceled.
```

[3] The main idea here would be to pay rides on the Tesla Network in **KWH.2**
instead of any specific currency. Teslas could be configured to trust the Tesla
Network solely **modelx_AX7GD@tesla.com[KWH.2]/network@tesla.com[KWH.2] 1/1**
with the Tesla Network trusting the Superchargers Network with
**network@tesla.com[KWH.2]/superchargers@tesla.com[KWH.2] 11/10** such that
Tesla would get a cut of each transactions (10% here). Reciprocally the
Superchargers Network could trust the Tesla Network at a flat rate
**network@tesla.com[KWH.2]/superchargers@tesla.com[KWH.2] 1/1** to let Teslas
pay for their recharge as they earn **KWH.2**.  Car-owners would also trust
**superchargers@tesla.com[KWH.2]** using it to "pay" at Supercharger stations
when they ride their own Tesla which would in turn enable Teslas to pay for
charges at home stations. From there Tesla could set prices for their
Supercharger in each currency by trusting local nodes at various rates, or even
local power plants for **KWH.2** directly.

[4] Online merchants could use Settle to issue store credit that is easily
usable at other places. This could potentially solve the problem of unlinked
refunds in e-commerce while ensuring merchants that these credits would, by
construction, be eventually used at their store.

[5] Lets users deposit currencies (fiat or cryptocurrencies) in exchange for
trust on the network, and get that currency deposited back to their account by
paying the gateway on the network: by depositing $20 at acme.com (using a card
payment or transfer), I would receive a trustline from **gw@acme.com[USD.2]**
to **spolu@m.settle.network[USD.2]**. Conversely, when paying **gw@acme.com
USD.2 100**, it would automatically deposit back the associated funds to my
bank account.

[6]
[settle-public@googlegroups.com](https://groups.google.com/d/forum/settle-public)


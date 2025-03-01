# Contributing to Quorum Key Manager

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

Welcome to the Quorum Key Manager repositories! The following is a set of guidelines for contributing to this repo and its packages. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

#### Table Of Contents

[Code of Conduct](#code-of-conduct)

[I just have a quick question](#i-just-have-a-quick-question)

[How to contribute](#how-to-contribute)

- [Your first code contribution](#your-first-code-contribution)
- [Reporting bugs](#reporting-bugs)

[Style Guides](#style-guides)

- [GO Style Guide](#GO-style-guide)
- [Git Commit Messages & Pull Request Messages](#git-commit-messages--pull-request-messages)

## Code of Conduct

This project and everyone participating in it is governed by the [Quorum Key Manager Code of Conduct](CODE-OF-CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [private-quorum@consensys.net].

## I just have a quick question

Send us an email at [quorum-key-manager@consensys.net](mailto:quorum-key-manager@consensys.net) and that's the fastest way to get an answer.

## How To Contribute

### Your first code contribution

Start by looking through the 'good first issue' and 'help wanted' issues:

- [Good First Issue][search-label-good-first-issue] - issues which should only require a few lines of code, and a test or two.
- [Help wanted issues][search-label-help-wanted] - issues which are a bit more involved than `good first issue` issues.

### Local Development

The codebase is maintained using the "_contributor workflow_" where everyone without exception contributes patch proposals using "_pull-requests_". This facilitates social contribution, easy testing and peer review.

To contribute a patch, the workflow is as follows:

- Fork repository
- Create topic branch
- Commit patch
- Create pull-request, adhering to the coding conventions herein set forth

In general a commit serves a single purpose and diffs should be easily comprehensible. For this reason do not mix any formatting fixes or code moves with actual code changes.

### Automated Test coverage

All code submissions must be accompanied by appropriate automated tests. The goal is to provide confidence in the code’s robustness, while avoiding redundant tests.

### Pull Requests

The process described here has several goals:

- Maintain Product quality
- Fix problems that are important to users
- Engage the community in working toward the best possible product
- Enable a sustainable system for maintainers to review contributions
- Further explanation on PR & commit messages can be found in this post: [How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/).

Please follow these steps to have your contribution considered by the approvers:

1. Complete the CLA, as described in [CLA.md]
2. Follow all instructions in [PULL-REQUEST-TEMPLATE.md](.github/pull_request_template.md)
3. Include appropriate test coverage. Testing is 100% automated. There is no such thing as a manual test.
4. Follow the [Style Guides](#style-guides)
5. After you submit your pull request, verify that all [status checks](https://help.github.com/articles/about-status-checks/) are passing <details><summary>What if the status checks are failing?</summary>If a status check is failing, and you believe that the failure is unrelated to your change, please leave a comment on the pull request explaining why you believe the failure is unrelated. A maintainer will re-run the status check for you. If we conclude that the failure was a false positive, then we will open an issue to track that problem with our status check suite.</details>

While the prerequisites above must be satisfied prior to having your pull request reviewed, the reviewer(s) may ask you to complete additional design work, tests, or other changes before your pull request can be ultimately accepted. Please refer to [Code Reviews].

## Reporting Bugs

Following these guidelines helps maintainers and the community understand your report, reproduce the behavior, and find related reports.

Explain the problem and include additional details to help maintainers reproduce the problem:

- **Use a clear and descriptive title** for the issue to identify the problem.
- **Describe the exact steps which reproduce the problem** in as many details as possible. For example, start by explaining how you started Quorum Key Manager, e.g. which command exactly you used in the terminal, or how you started it otherwise.
- **Provide specific examples to demonstrate the steps**. Include links to files or GitHub projects, or copy/pasteable snippets, which you use in those examples. If you're providing snippets in the issue, use [Markdown code blocks](https://help.github.com/articles/markdown-basics/#multiple-lines).
- **Describe the behavior you observed after following the steps** and point out what exactly is the problem with that behavior.
- **Explain which behavior you expected to see instead and why.**
- **Include screenshots** which show you following the described steps and clearly demonstrate the problem.

Provide more context by answering these questions:

- **Did the problem start happening recently** (e.g. after updating to a new version of the software) or was this always a problem?
- If the problem started happening recently, **can you reproduce the problem in an older version of the software?** What's the most recent version in which the problem doesn't happen?
- **Can you reliably reproduce the issue?** If not, provide details about how often the problem happens and under which conditions it normally happens.

Include details about your configuration and environment:

- **What OS & Version are you running?**
  - **For Linux - What kernel are you running?** You can get the exact version by running `uname -a` in your terminal.
- **Are you running in a virtual machine?** If so, which VM software are you using and which operating systems and versions are used for the host and the guest?
- **Are you running in a docker container?** If so, what version of docker?
- **Are you running in a Cloud?** If so, which one, and what type/size of VM is it?

# Style Guides

## GO Code Style Guide

To reformat code, run:

```
make lint
```

Code style will be checked automatically during a build.

## Git Commit Messages & Pull Request Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Provide a summary on the first line with more details on additional lines as needed
- Reference issues and pull requests liberally

[private-quorum@consensys.net]: mailto:private-quorum@consensys.net
[cla.md]: /CLA.md

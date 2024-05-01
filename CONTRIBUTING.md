# Welcome to the Contributor's Guide

Thank you for dedicating your time to contribute to this project. Your contributions are greatly valued! 😺

This guide serves to facilitate your journey into the contribution process.

## Contributing Opportunities

This project encourages contributions from the community. You can contribute in the following areas:

### Feedback

Contribute to the project's enhancement by sharing your insights and ideas. Open a [feature request issue](https://github.com/MesaTechLabs/kitten/issues/new?assignees=&labels=enhancement%2Ctriage&projects=&template=feature-request.yaml&title=%5BFeature%5D%3A+) to initiate dialogue.

### Documentation

Participate in refining project documentation. Whether rectifying a typo, fixing a broken link, or envisioning broader enhancements, your contributions are valued. For minor amendments, feel free to submit a PR directly. For substantial changes, start a discussion by opening an issue.

### Bug Discovery

Contribute to project stability by identifying and reporting bugs. Before flagging an issue, [ensure it hasn't already been addressed](https://github.com/MesaTechLabs/kitten/issues). If untouched, seize the opportunity to address it yourself or leave it open for others to tackle.

## Contribution Workflow

To ensure smooth collaboration and efficient handling of contributions, the contribution process is outlined as follows:

1. **Issue Initiation**:
   - Users initiate issues for various purposes such as reporting bugs, requesting features, proposing changes, or asking questions.
2. **Issue Review and Labeling**:
   - Maintainers review incoming issues and categorize them with appropriate status labels:
     - `triage`: Issues requiring further evaluation or assignment.
     - `help wanted`: Issues ready for work, awaiting volunteer contribution.
     - `duplicate`: Issues that duplicate another and require consolidation.
     - `blocked`: Issues impeded by dependencies or external factors.
     - `needs revision`: Issues lacking necessary information or clarity.
     - `under review`: Issues pending further discussion or decision-making.
     - `wontfix`: Issues falling outside the project's scope and not planned for resolution.
     - `BREAKING CHANGE`:  Issues that introduce significant alterations to the project's codebase or functionality, potentially breaking backward compatibility or requiring substantial reworking of existing implementations.
3. **Volunteer Engagement**:
   - Issue authors or other volunteers express interest in resolving issues by offering solutions or volunteering for ownership.
4. **Assignment**:
   - Maintainers assign issues to volunteers, designating them as the issue's "owner".
5. **Pull Request Submission**:
   - The issue owner submits a pull request containing proposed changes to address the identified issue.
6. **Review and Merge**:
   - Maintainers review the pull request, provide feedback if necessary, and merge it into the main repository upon approval.

> [!IMPORTANT]
> Before submitting a pull request, except for minor documentation/typo fixes, please ensure you've opened an issue and had it assigned to you. This precaution prevents wasted effort on submissions that might not be merged or accepted.

> [!NOTE]
> To volunteer to resolve someone else's issue, kindly comment on the issue expressing your interest in working on it. Include a brief description of your proposed approach if it's not already evident from the issue description and the issue will be assigned to you.

## Good First Issues

If you're new to the codebase and seeking a starting point, consider checking out issues labeled [good first issue](https://github.com/MesaTechLabs/kitten/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22+-label%3A%22blocked+by+upstream%22) since they should be relatively straightforward to work on. Before diving in, ensure there's no existing PR for the issue and that it hasn't been assigned to anyone yet. Once you've identified an issue you'd like to tackle, please notify the maintainers by commenting on the issue. This ensures proper coordination and prevents duplicate efforts.

## Changesets Workflow

This project uses [Changesets](https://github.com/changesets/changesets) for package versioning. On this project, it's mainly used to automate the GitHub release through the [CI GitHub Action](.github/workflows/ci.yaml) and update the [CHANGELOG.md](./CHANGELOG.md) automatically.

The workflow looks as follows:

1. **Implement Feature Adjustments**:
   - Make all necessary adjustments to the feature you are currently working on.

2. **Generate Changesets**:
   - Use `bun changeset` to generate a changeset, following the provided instructions. Remember to classify changesets introducing new features as `minor`, while those addressing bugs should be marked as `patch`.

3. **Commit with Changeset**:
   - Make a commit that encompasses all relevant adjustments, ensuring to include the generated changeset.

4. **Submit Pull Request**:
   - Once the commit is ready, submit a pull request to integrate your changes.

> [!IMPORTANT]
> Only generate changesets with `major` if the issue is labeled with `BREAKING CHANGE`.

## License

By contributing to this project, you agree to license your contributions under the project's  [license](./LICENSE).

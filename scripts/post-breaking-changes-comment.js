/**
 * Script to post breaking changes analysis as a comment on the PR
 * Used by GitHub Actions workflow
 */

module.exports = async ({ github, context, core }) => {
  // Get values from environment variables
  const breakingChangesDetected = process.env.BREAKING_CHANGES_DETECTED === 'true';
  const breakingChangesJson = JSON.parse(process.env.BREAKING_CHANGES_JSON || '{}');

  // Prepare comment content
  let commentBody = '## Breaking Changes Analysis\n\n';

  if (breakingChangesDetected) {
    commentBody += '⚠️ **BREAKING CHANGES DETECTED** ⚠️\n\n';
  } else {
    commentBody += '✅ **No breaking changes detected**\n\n';
  }

  // If we have result data
  if (breakingChangesJson && Object.keys(breakingChangesJson).length > 0) {
    // Add removed resources if any
    if (breakingChangesJson.removed_resources && breakingChangesJson.removed_resources.length > 0) {
      commentBody += '### Removed Resources\n';
      breakingChangesJson.removed_resources.forEach((resource) => {
        commentBody += `- \`${resource}\`\n`;
      });
      commentBody += '\n';
    }

    // Add changed resources if any
    if (
      breakingChangesJson.changed_resources &&
      Object.keys(breakingChangesJson.changed_resources).length > 0
    ) {
      commentBody += '### Changed Resources\n';

      Object.entries(breakingChangesJson.changed_resources).forEach(([resource, changes]) => {
        commentBody += `#### \`${resource}\`\n`;

        // Breaking changes
        if (changes.breaking_changes && changes.breaking_changes.length > 0) {
          commentBody += '**Breaking Changes:**\n';
          changes.breaking_changes.forEach((change) => {
            commentBody += `- \`${change.name}\`: ${change.type}`;
            if (change.description) {
              commentBody += ` (${change.description})`;
            }
            commentBody += '\n';
          });
          commentBody += '\n';
        }

        // Non-breaking changes
        if (changes.non_breaking_changes && changes.non_breaking_changes.length > 0) {
          commentBody += '**Non-Breaking Changes:**\n';
          changes.non_breaking_changes.forEach((change) => {
            commentBody += `- \`${change.name}\`: ${change.type}`;
            if (change.description) {
              commentBody += ` (${change.description})`;
            }
            commentBody += '\n';
          });
          commentBody += '\n';
        }
      });
    }

    // Add new resources if any
    if (breakingChangesJson.added_resources && breakingChangesJson.added_resources.length > 0) {
      commentBody += '### Added Resources\n';
      breakingChangesJson.added_resources.forEach((resource) => {
        commentBody += `- \`${resource}\`\n`;
      });
      commentBody += '\n';
    }

    // Add new providers if any
    if (breakingChangesJson.added_providers && breakingChangesJson.added_providers.length > 0) {
      commentBody += '### Added Providers\n';
      breakingChangesJson.added_providers.forEach((provider) => {
        commentBody += `- \`${provider}\`\n`;
      });
      commentBody += '\n';
    }
  }

  commentBody += '---\n';
  commentBody += 'For more details, see the CI logs.';

  // Post comment to PR
  const { data: comments } = await github.rest.issues.listComments({
    ...context.repo,
    issue_number: context.issue.number,
  });

  // Check for existing breaking changes comment to update instead of creating a new one
  const breakingChangesComment = comments.find(
    (comment) =>
      comment.user.login === 'github-actions[bot]' && comment.body.includes('Breaking Changes Analysis'),
  );

  if (breakingChangesComment) {
    await github.rest.issues.updateComment({
      ...context.repo,
      comment_id: breakingChangesComment.id,
      body: commentBody,
    });
    console.log('Updated existing breaking changes comment');
  } else {
    await github.rest.issues.createComment({
      ...context.repo,
      issue_number: context.issue.number,
      body: commentBody,
    });
    console.log('Created new breaking changes comment');
  }
};

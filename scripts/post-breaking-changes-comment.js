/**
 * Script to post breaking changes analysis as a comment on the PR
 * Used by GitHub Actions workflow
 */

module.exports = async ({ github, context, core }) => {
  try {
    // Get values from environment variables
    const breakingChangesDetected = process.env.BREAKING_CHANGES_DETECTED === 'true';
    const rawJson = process.env.BREAKING_CHANGES_JSON || '{}';
    
    // Prepare comment content
    let commentBody = '## Breaking Changes Analysis\n\n';

    if (breakingChangesDetected) {
      commentBody += '⚠️ **BREAKING CHANGES DETECTED** ⚠️\n\n';
    } else {
      commentBody += '✅ **No breaking changes detected**\n\n';
    }

    // Add raw JSON output
    commentBody += '```json\n';
    commentBody += rawJson;
    commentBody += '\n```\n\n';

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
        comment.user.login === 'github-actions[bot]' && comment.body.includes('Breaking Changes Analysis')
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
  } catch (error) {
    console.error('Error posting comment:', error);
    throw error; // Re-throw to ensure the workflow fails
  }
};

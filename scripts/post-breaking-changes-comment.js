/**
 * Script to post breaking changes analysis as a comment on the PR
 * Used by GitHub Actions workflow
 */

module.exports = async ({ github, context, core }) => {
  console.log('DEBUG: Script started');
  try {
    console.log('DEBUG: Node version:', process.version);
    console.log('DEBUG: Environment variables:', Object.keys(process.env));
    
    // Get values from environment variables
    console.log('DEBUG: Reading environment variables');
    console.log('DEBUG: BREAKING_CHANGES_DETECTED =', process.env.BREAKING_CHANGES_DETECTED);
    console.log('DEBUG: BREAKING_CHANGES_JSON length =', (process.env.BREAKING_CHANGES_JSON || '').length);
    console.log('DEBUG: BREAKING_CHANGES_JSON substring =', (process.env.BREAKING_CHANGES_JSON || '').substring(0, 100) + '...');
    
    const breakingChangesDetected = process.env.BREAKING_CHANGES_DETECTED === 'true';
    console.log('DEBUG: breakingChangesDetected =', breakingChangesDetected);
    
    let breakingChangesJson;
    try {
      console.log('DEBUG: Parsing JSON');
      breakingChangesJson = JSON.parse(process.env.BREAKING_CHANGES_JSON || '{}');
      console.log('DEBUG: JSON parsed successfully');
    } catch (parseError) {
      console.error('DEBUG: Error parsing JSON:', parseError);
      console.log('DEBUG: JSON content that failed parsing:', process.env.BREAKING_CHANGES_JSON);
      breakingChangesJson = {};
    }
    
    console.log('DEBUG: Creating comment body');
    // Prepare comment content
    let commentBody = '## Breaking Changes Analysis\n\n';

    if (breakingChangesDetected) {
      console.log('DEBUG: Adding breaking changes detected message');
      commentBody += '⚠️ **BREAKING CHANGES DETECTED** ⚠️\n\n';
    } else {
      console.log('DEBUG: Adding no breaking changes message');
      commentBody += '✅ **No breaking changes detected**\n\n';
    }

    // If we have result data
    console.log('DEBUG: Checking breakingChangesJson', typeof breakingChangesJson, Object.keys(breakingChangesJson).length > 0);
    if (breakingChangesJson && Object.keys(breakingChangesJson).length > 0) {
      console.log('DEBUG: Processing JSON data');
      
      // Add removed resources if any
      console.log('DEBUG: Checking removed_resources');
      if (breakingChangesJson.removed_resources && breakingChangesJson.removed_resources.length > 0) {
        console.log('DEBUG: Adding removed resources section');
        commentBody += '### Removed Resources\n';
        breakingChangesJson.removed_resources.forEach((resource) => {
          console.log('DEBUG: Adding removed resource:', resource);
          commentBody += `- \`${resource}\`\n`;
        });
        commentBody += '\n';
      }

      // Add changed resources if any
      console.log('DEBUG: Checking changed_resources');
      console.log('DEBUG: changed_resources =', breakingChangesJson.changed_resources);
      console.log('DEBUG: changed_resources type =', typeof breakingChangesJson.changed_resources);
      if (
        breakingChangesJson.changed_resources &&
        Object.keys(breakingChangesJson.changed_resources).length > 0
      ) {
        console.log('DEBUG: Adding changed resources section');
        commentBody += '### Changed Resources\n';

        Object.entries(breakingChangesJson.changed_resources).forEach(([resource, changes]) => {
          console.log('DEBUG: Processing resource:', resource);
          console.log('DEBUG: Resource changes:', changes);
          commentBody += `#### \`${resource}\`\n`;

          // Breaking changes
          console.log('DEBUG: Checking breaking_changes for resource:', resource);
          if (changes.breaking_changes && changes.breaking_changes.length > 0) {
            console.log('DEBUG: Adding breaking_changes');
            commentBody += '**Breaking Changes:**\n';
            changes.breaking_changes.forEach((change) => {
              console.log('DEBUG: Processing breaking change:', change);
              commentBody += `- \`${change.name}\`: ${change.type}`;
              if (change.description) {
                commentBody += ` (${change.description})`;
              }
              commentBody += '\n';
            });
            commentBody += '\n';
          }

          // Non-breaking changes
          console.log('DEBUG: Checking non_breaking_changes for resource:', resource);
          if (changes.non_breaking_changes && changes.non_breaking_changes.length > 0) {
            console.log('DEBUG: Adding non_breaking_changes');
            commentBody += '**Non-Breaking Changes:**\n';
            changes.non_breaking_changes.forEach((change) => {
              console.log('DEBUG: Processing non-breaking change:', change);
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
      console.log('DEBUG: Checking added_resources');
      if (breakingChangesJson.added_resources && breakingChangesJson.added_resources.length > 0) {
        console.log('DEBUG: Adding added resources section');
        commentBody += '### Added Resources\n';
        breakingChangesJson.added_resources.forEach((resource) => {
          console.log('DEBUG: Adding added resource:', resource);
          commentBody += `- \`${resource}\`\n`;
        });
        commentBody += '\n';
      }

      // Add new providers if any
      console.log('DEBUG: Checking added_providers');
      if (breakingChangesJson.added_providers && breakingChangesJson.added_providers.length > 0) {
        console.log('DEBUG: Adding added providers section');
        commentBody += '### Added Providers\n';
        breakingChangesJson.added_providers.forEach((provider) => {
          console.log('DEBUG: Adding added provider:', provider);
          commentBody += `- \`${provider}\`\n`;
        });
        commentBody += '\n';
      }
    }

    commentBody += '---\n';
    commentBody += 'For more details, see the CI logs.';

    console.log('DEBUG: Comment body generated');
    console.log('DEBUG: Comment body length =', commentBody.length);

    // Post comment to PR
    console.log('DEBUG: Listing existing comments');
    console.log('DEBUG: context.repo =', context.repo);
    console.log('DEBUG: context.issue.number =', context.issue.number);
    
    const { data: comments } = await github.rest.issues.listComments({
      ...context.repo,
      issue_number: context.issue.number,
    });
    console.log('DEBUG: Found', comments.length, 'existing comments');

    // Check for existing breaking changes comment to update instead of creating a new one
    console.log('DEBUG: Looking for existing breaking changes comment');
    const breakingChangesComment = comments.find(
      (comment) => {
        console.log('DEBUG: Checking comment from user:', comment.user.login);
        return comment.user.login === 'github-actions[bot]' && comment.body.includes('Breaking Changes Analysis');
      }
    );

    if (breakingChangesComment) {
      console.log('DEBUG: Found existing comment, will update. ID =', breakingChangesComment.id);
      await github.rest.issues.updateComment({
        ...context.repo,
        comment_id: breakingChangesComment.id,
        body: commentBody,
      });
      console.log('DEBUG: Updated existing breaking changes comment');
    } else {
      console.log('DEBUG: No existing comment found, will create new one');
      await github.rest.issues.createComment({
        ...context.repo,
        issue_number: context.issue.number,
        body: commentBody,
      });
      console.log('DEBUG: Created new breaking changes comment');
    }
    
    console.log('DEBUG: Script completed successfully');
  } catch (error) {
    console.error('DEBUG: Caught error in main function:', error);
    console.error('DEBUG: Error name:', error.name);
    console.error('DEBUG: Error message:', error.message);
    console.error('DEBUG: Error stack trace:', error.stack);
    throw error; // Re-throw to ensure the workflow fails
  }
};

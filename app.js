fetch('http://localhost:8080/comments')
			.then(response => response.json())
			.then(data => {
				// dynamically create HTML elements to display the JSON data
				let commentsDiv = document.getElementById('comments');
				for (let comment of data) {
					let commentDiv = document.createElement('div');
					commentDiv.innerHTML = `<h2>${comment.text}</h2>`;
					if (comment.child_comments) {
						for (let childComment of comment.child_comments) {
							let childCommentDiv = document.createElement('div');
							childCommentDiv.innerHTML = `<p>${childComment.text}</p>`;
							if (childComment.child_comments) {
								for (let grandchildComment of childComment.child_comments) {
									let grandchildCommentDiv = document.createElement('div');
									grandchildCommentDiv.innerHTML = `<p>${grandchildComment.text}</p>`;
									childCommentDiv.appendChild(grandchildCommentDiv);
								}
							}
							commentDiv.appendChild(childCommentDiv);
						}
					}
					commentsDiv.appendChild(commentDiv);
				}
			})
			.catch(error => console.error(error));
var ghpages = require("gh-pages");

ghpages.publish("/gh-pages/dist", {
    branch: "gh-pages",
    user: {
	name: "task4233",
	email: "mail@task4233.dev"
    },
    message: "Auto deploy for goDoc page [ci skip]"
}, function(err) {
    if (err) {
	console.log(err);
	process.exit(1);
    } else {
	console.log("Successfully updated!");
    }
});

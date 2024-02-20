let password = "";
let passwordFailCounter = 0;

/**
 * Validates the password entered when first a user first enters the site
 */
async function validatePassword() {
    if (passwordFailCounter >= 5) {
        window.location.href = "https://minecraft.net";
        return;
    }

    password = prompt("Password");

    if (password !== "") {
        const validation = await fetch("api/validate", {
            method: "POST",
            headers: {
                "X-Api-Key": password,
            }
        });

        const response = await validation.json();
        if (response.success !== true) {
            passwordFailCounter++;
            alert("Wrong password");
            await validatePassword();
        }
    }
}

/**
 * Function to handle form submission
 * @param {Event} e 
 * @returns void
 */
async function whitelistUser(e) {
    e.preventDefault();

    const form = document.getElementById("whitelistForm");
    const formData = new FormData(form);

    const username = formData.get("username");
    if (username === undefined || username === "") {
        alert("Please enter a valid username");
        return false;
    }

    if (!validateUsername(username)) {
        alert("Please enter a valid username");
        return false;
    }

    try {
        const request = await fetch(
            "api/whitelist/add",
            {
                method: "POST",
                body: JSON.stringify({
                    username
                }),
                headers: {
                    "X-Api-Key": password,
                    "Content-Type": "application/json",
                },
            },
        );

        const response = await request.json();
    
        if (response.success !== true) {
            alert(response.message);
            return false;
        }

        alert(response.message);
        form.reset();
    } catch (err) {
        console.log(`Error ${err.message}`);
        alert("There was an error adding to the whitelist");
    }
    
}

/**
 * Validates if the username is invalid by splitting by string
 * @param {string} username 
 * @returns 
 */
function validateUsername(username) {
    return username.split(" ").length === 1;
}

// Entrypoint
validatePassword();
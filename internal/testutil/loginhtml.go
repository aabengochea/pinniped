// Copyright 2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"fmt"

	"go.pinniped.dev/internal/here"
)

func ExpectedLoginPageHTML(wantCSS, wantIDPName, wantPostPath, wantEncodedState, wantAlert string) string {
	alertHTML := ""
	if wantAlert != "" {
		alertHTML = fmt.Sprintf("\n"+
			"    <div class=\"form-field\">\n"+
			"        <span class=\"alert\" role=\"alert\" aria-label=\"login error message\" id=\"alert\">%s</span>\n"+
			"    </div>\n    ",
			wantAlert,
		)
	}

	return here.Docf(`<!DOCTYPE html>
        <html lang="en">
        <head>
            <title>Pinniped Login</title>
            <meta charset="UTF-8">
            <style>%s</style>
            <link href="data:image/x-icon;base64,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"
                  rel="icon" type="image/x-icon"/>
        </head>
        <body>
        <div class="box" aria-label="login form" role="main">
            <div class="form-field">
                <h1>Log in to %s</h1>
            </div>
            %s
            <form action="%s" method="post">
                <input type="hidden" name="state" id="state" value="%s">
                <div class="form-field">
                    <label for="username"><span class="hidden" aria-hidden="true">Username</span></label>
                    <input type="text" name="username" id="username"
                           autocomplete="username" placeholder="Username" required>
                </div>
                <div class="form-field">
                    <label for="password"><span class="hidden" aria-hidden="true">Password</span></label>
                    <input type="password" name="password" id="password"
                           autocomplete="current-password" placeholder="Password" required>
                </div>
                <div class="form-field">
                    <input type="submit" name="submit" id="submit" value="Log in"/>
                </div>
            </form>
        </div>
        </body>
        </html>
	`,
		wantCSS,
		wantIDPName,
		alertHTML,
		wantPostPath,
		wantEncodedState,
	)
}

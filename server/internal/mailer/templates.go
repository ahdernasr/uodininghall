package mailer

func GetWelcomeTemplate() string {
	welcomeTemplate := `
<!DOCTYPE html>
<html>
<head>
	<title>Subcription Confirmation</title>
<style>
body {
	font-family: 'Arial', sans-serif;
	margin: 0;
	padding: 0;
}
.container {
	background-color: #ffffff;
	margin: 20px auto;
	padding: 20px;
	max-width: 600px;
	border-radius: 10px;
	border: 1px solid black;
}
h1 {
	text-align: center;
	color: #ffffff; /* White text for the main title */
	background-color: #8F001A; /* University of Ottawa's Garnet */
	padding: 10px;
	border-radius: 10px;
}

/* Dark mode styles */
@media (prefers-color-scheme: dark) {
	body {
		background-color: #000000; /* Dark background for the main content */
		color: #e0e0e0; /* Light grey text for dark mode */
	}
	.container {
		background-color: #333333;
		/* box-shadow: 0 0 10px rgba(255, 255, 255, 0.1); */
		border: 1px solid black;
	}
	h1 {
		color: #000000; /* Dark text for the main title */
		background-color: #8F001A; /* University of Ottawa's Garnet */
	}
}
</style>
</head>
<body>
	<div class="container">
		You have subscribed to UO Dining Hall.
	</div>
</body>
</html>`

	return welcomeTemplate
}

func getMenuTemplate() string {

	menuTemplate := `
<!DOCTYPE html>
<html>
<head>
	<title>Today's Menu</title>
<style>
body {
	font-family: 'Arial', sans-serif;
	margin: 0;
	padding: 0;
}
.container {
	background-color: #ffffff;
	margin: 20px auto;
	padding: 20px;
	max-width: 600px;
	border-radius: 10px;
	border: 1px solid black;
}
h1 {
	text-align: center;
	color: #ffffff; /* White text for the main title */
	background-color: #8F001A; /* University of Ottawa's Garnet */
	padding: 10px;
	border-radius: 10px;
}
h2 {
	color: #8F001A; /* University of Ottawa's Garnet for section titles */
	border-bottom: 1px solid #8F001A;
	padding-bottom: 5px;
}
h3 {
	color: #80746C; /* Grey for subheadings */
	margin-top: 20px;
}
p {
	color: #80746C; /* Grey for regular text */
	line-height: 1.5;
}
p strong {
	color: #8F001A; /* University of Ottawa's Garnet for emphasis */
}

/* Dark mode styles */
@media (prefers-color-scheme: dark) {
	body {
		background-color: #000000; /* Dark background for the main content */
		color: #e0e0e0; /* Light grey text for dark mode */
	}
	.container {
		background-color: #333333;
		/* box-shadow: 0 0 10px rgba(255, 255, 255, 0.1); */
		border: 1px solid black;
	}
	h1 {
		color: #000000; /* Dark text for the main title */
		background-color: #8F001A; /* University of Ottawa's Garnet */
	}
	h2 {
		color: #e0e0e0; /* Light grey for section titles */
		border-bottom: 1px solid #e0e0e0;
	}
	h3 {
		color: #c0c0c0; /* Lighter grey for subheadings */
	}
	p {
		color: #e0e0e0; /* Light grey for regular text */
	}
	p strong {
		color: #ffcccc; /* Lighter Garnet for emphasis */
	}
}
</style>
</head>
<body>
	<div class="container">
		<h1>Today's Menu</h1>
		<h2>Grill</h2>
		{{if .Grill.Breakfast}}
		<h3>Breakfast</h3>
		{{range .Grill.Breakfast}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .Grill.Lunch}}
		<h3>Lunch</h3>
		{{range .Grill.Lunch}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .Grill.Dinner}}
		<h3>Dinner</h3>
		{{range .Grill.Dinner}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .Grill.Other}}
		<h3>Special</h3>
		{{range .Grill.Other}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		<h2>Mind Body Soul</h2>
		{{if .MindBodySoul.Breakfast}}
		<h3>Breakfast</h3>
		{{range .MindBodySoul.Breakfast}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .MindBodySoul.Lunch}}
		<h3>Lunch</h3>
		{{range .MindBodySoul.Lunch}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .MindBodySoul.Dinner}}
		<h3>Dinner</h3>
		{{range .MindBodySoul.Dinner}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .MindBodySoul.Other}}
		<h3>Special</h3>
		{{range .MindBodySoul.Other}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		<h2>Plant Based</h2>
		{{if .PlantBase.Other}}
		{{range .PlantBase.Other}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		<h2>Service Minute</h2>
		{{if .ServiceMinute.Breakfast}}
		<h3>Breakfast</h3>
		{{range .ServiceMinute.Breakfast}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .ServiceMinute.Lunch}}
		<h3>Lunch</h3>
		{{range .ServiceMinute.Lunch}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .ServiceMinute.Dinner}}
		<h3>Dinner</h3>
		{{range .ServiceMinute.Dinner}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .ServiceMinute.Other}}
		<h3>Special</h3>
		{{range .ServiceMinute.Other}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		<h2>Trattoria</h2>
		{{if .Trattoria.Breakfast}}
		<h3>Breakfast</h3>
		{{range .Trattoria.Breakfast}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .Trattoria.Lunch}}
		<h3>Lunch</h3>
		{{range .Trattoria.Lunch}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .Trattoria.Dinner}}
		<h3>Dinner</h3>
		{{range .Trattoria.Dinner}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .Trattoria.Other}}
		<h3>Special</h3>
		{{range .Trattoria.Other}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		<h2>World Flavours</h2>
		{{if .WorldFlavours.Breakfast}}
		<h3>Breakfast</h3>
		{{range .WorldFlavours.Breakfast}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .WorldFlavours.Lunch}}
		<h3>Lunch</h3>
		{{range .WorldFlavours.Lunch}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .WorldFlavours.Dinner}}
		<h3>Dinner</h3>
		{{range .WorldFlavours.Dinner}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
		{{if .WorldFlavours.Other}}
		<h3>Special</h3>
		{{range .WorldFlavours.Other}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		{{end}}
	</div>
</body>
</html>`
	return menuTemplate
}

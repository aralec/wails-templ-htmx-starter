# README


## What is this?

This project is a hobby, because I like coding in Go and I want to progress on this language. It may not be written according to the best standards - it is a work in progress.

I hope to use it as a starting bloc of any future side project.

The goal of this repo is to provide the template of a web application written in Go and Templ + HTMX. This template would not embed any framework or CSS library, to remain as neutral as possible.

The code of this application, here a Counter for the example, is embedded in the /src folder; the rest is necessary for the operation of Wails.

The idea is to have in /src an application that could be extracted, almost "as is". The /internal folder carries the backend logic of the application, organized according to a hexagonal architecture. The /external folder carries the front-end code, here written using Templ.

In /src, the code also follows a "plug-in" philosophy, because the internal folder must provide a reusable application core if desired with a front-end technology other than Templ; hence the use of a hexagonal architecture.

## Credits

This repo is based on https://github.com/PylotLight/wails-htmx-templ-template

And modified to follow hexagonal architecture recommendations that I've read online : 
https://www.golinuxcloud.com/hexagonal-architectural-golang/

The Counter code is an adaptation of https://templ.guide/server-side-rendering/example-counter-application

## About

This template uses a unique combination of pure htmx for interactivity plus Go templates for creating components and forms, also included:
- Uses HTMX for MPA style interactivity on a single page as per SPA.
- Added Chi middleware for handling HTMX calls in an easy to maintain routing configuration.
- Built-in version display linked to version variable from main which can be updated on app build for CICD and/or during runtime.
- Scripts configured to use the Bun runtime to launch Vite. (Make sure you have bun installed first)
- To switch back to npm instead of bun, edit wails.json and package.json (or use NPM under the @npm tag)
- Also using https://templ.guide/ for components and templates use "templ generate" to update templ files. 

## Initial Setup Instructions
- Install Bun (or use NPM under the @npm tag)
- Change go.mod module
- Change app.go components package import
- **For Linux build tag webkit2_40 is required e.g -tags webkit2_40**

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `bun run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

To build a redistributable, production mode package, use `wails build`.

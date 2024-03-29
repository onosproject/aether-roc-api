<!--
~ SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
~
~ SPDX-License-Identifier: Apache-2.0
-->
<!DOCTYPE html>
<html>
<head>
  <title>{{.Description}}</title>
  <!-- needed for adaptive design -->
  <meta charset="utf-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

  <!--
  ReDoc doesn't change outer page styles
  -->
  <style>
    body {
      margin: 0;
      padding: 0;
    }
  </style>
</head>
<body>
<redoc spec-url='{{.File}}'></redoc>
<script src="redoc.standalone.js"> </script>
</body>
</html>

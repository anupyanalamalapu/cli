import React from 'react';
import PropTypes from 'prop-types';
import { Header, Jumbotron } from 'watson-react-components';

export default function Layout(props) {
  return (
    <html lang="en">
      <head>
        <title>Afya</title>
        <meta charSet="utf-8" />
        <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/images/favicon.ico" type="image/x-icon" />
        <link rel="stylesheet" href="/css/watson-react-components.min.css" />
        <link rel="stylesheet" href="/css/style.css" />
      </head>
      <body>
        <Header
          mainBreadcrumbs="Afya"
          mainBreadcrumbsUrl="."
        />
        <div id="root">
          {props.children}
        </div>
        <script type="text/javascript" src="scripts/bundle.js" />
        { props.bluemixAnalytics ? <script type="text/javascript" src="scripts/analytics.js" /> : null }
      </body>
    </html>
  );
}

Layout.propTypes = {
  children: PropTypes.object.isRequired, // eslint-disable-line
  bluemixAnalytics: PropTypes.bool.isRequired,
};

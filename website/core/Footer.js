/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const React = require('react');

class Footer extends React.Component {
  docUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    const docsUrl = this.props.config.docsUrl;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    return `${baseUrl}${docsPart}${langPart}${doc}`;
  }

  pageUrl(doc, language) {
    const baseUrl = this.props.config.baseUrl;
    return baseUrl + (language ? `${language}/` : '') + doc;
  }

  render() {
    return (
      <footer className="nav-footer" id="footer">
        <section className="sitemap">
          <a href={this.props.config.baseUrl} className="nav-home">
            {this.props.config.footerIcon && (
              <img
                src={this.props.config.baseUrl + this.props.config.footerIcon}
                alt={this.props.config.title}
                width="66"
                height="58"
              />
            )}
          </a>
          <div>
            <h5>Docs</h5>
            <a href={this.docUrl('Developer_Quickstart')}>
              Developer Quickstart
            </a>
            <a href={this.docUrl('AVM_Specification')}>
              AVM Specification
            </a>
          </div>
          <div>
            <h5>Community</h5>
            <a href="https://offchainlabs.com">Offchain Labs</a>
            <a href="https://medium.com/OffchainLabs">Blog</a>
            <a href="https://twitter.com/OffchainLabs">Twitter</a>
          </div>
          <div>
            <h5>More</h5>
            <a href="https://medium.com/OffchainLabs">Blog</a>
            <a href="https://github.com/OffchainLabs">GitHub</a>
          </div>
        </section>

        <a href="https://offchainlabs.com/">
          <img
            src={`${this.props.config.baseUrl}img/logo-white.png`}
            alt="Offchain Labs"
            width="150"
            height="125"
            className="footerLogo"
          />
        </a>
        <section className="copyright">{this.props.config.copyright}</section>
      </footer>
    );
  }
}

module.exports = Footer;

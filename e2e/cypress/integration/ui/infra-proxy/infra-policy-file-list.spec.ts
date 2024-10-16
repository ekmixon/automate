describe('infra policy file', () => {
  const now = Cypress.moment().format('MMDDYYhhmmss');
  const cypressPrefix = 'infra';
  let adminIdToken = '';
  const serverID = 'chef-server-dev-test';
  const serverName = 'chef server dev';
  const orgID = 'chef-org-dev';
  const orgName = '4thcoffee';
  const serverFQDN = 'ec2-34-219-25-251.us-west-2.compute.amazonaws.com';
  const serverIP = '34.219.25.251';
  const adminUser = 'chefadmin';
  const adminKey = Cypress.env('AUTOMATE_INFRA_ADMIN_KEY').replace(/\\n/g, '\n');
  const policyFileName = `${cypressPrefix}-policyFile-${now}-1`;

  before(() => {
    cy.adminLogin('/').then(() => {
      const admin = JSON.parse(<string>localStorage.getItem('chef-automate-user'));
      adminIdToken = admin.id_token;

      cy.request({
        auth: { bearer: adminIdToken },
        failOnStatusCode: false,
        method: 'POST',
        url: '/api/v0/infra/servers',
        body: {
          id: serverID,
          name: serverName,
          fqdn: serverFQDN,
          ip_address: serverIP
        }
      }).then((resp) => {
        if (resp.status === 200 && resp.body.ok === true) {
          return;
        } else {
          cy.request({
            auth: { bearer: adminIdToken },
            method: 'GET',
            url: `/api/v0/infra/servers/${serverID}`,
            body: {
              id: serverID
            }
          });
        }
      });

      cy.request({
        auth: { bearer: adminIdToken },
        failOnStatusCode: false,
        method: 'POST',
        url: `/api/v0/infra/servers/${serverID}/orgs`,
        body: {
          id: orgID,
          server_id: serverID,
          name: orgName,
          admin_user: adminUser,
          admin_key: adminKey
        }
      }).then((response) => {
        if (response.status === 200 && response.body.ok === true) {
          return;
        } else {
          cy.request({
            auth: { bearer: adminIdToken },
            method: 'GET',
            url: `/api/v0/infra/servers/${serverID}/orgs/${orgID}`,
            body: {
              id: orgID,
              server_id: serverID
            }
          });
        }
      });
      cy.visit(`/infrastructure/chef-servers/${serverID}/organizations/${orgID}`);
      cy.get('app-welcome-modal').invoke('hide');
    });
    cy.restoreStorage();
  });

  beforeEach(() => {
    cy.restoreStorage();
  });

  afterEach(() => {
    cy.saveStorage();
  });

  function getPolicyFile() {
    return cy.request({
      auth: { bearer: adminIdToken },
      failOnStatusCode: false,
      method: 'GET',
      url: `/api/v0/infra/servers/${serverID}/orgs/${orgID}/policyfiles`
    });
  }

  function checkResponse(response: any) {
    if (response.body.policies.length === 0) {
      cy.get('[data-cy=empty-list]').should('be.visible');
    } else {
      cy.get('[data-cy=policy-file-table-container] chef-th').contains('Name');
      cy.get('[data-cy=policy-file-table-container] chef-th').contains('Revision ID');
      return true;
    }
  }

  describe('infra policy file list page', () => {
    it('displays org details', () => {
      cy.get('.page-title').contains(orgName);
    });

    // policy file tabs specs
    it('can switch to policy file tab', () => {
      cy.get('.nav-tab').contains('Policyfiles').click();
    });

    it('can check if policy file has list or not', () => {
      getPolicyFile().then((response) => {
        checkResponse(response);
      });
    });

    context('can change page in policy file', () => {

      it('can change page and load data according to page', () => {
        getPolicyFile().then((response) => {
          if (checkResponse(response)) {
            if (cy.get('.policy-file-list-paging .page-picker-item').contains('3')) {
              cy.get('.policy-file-paging .page-picker-item').contains('3').click();
            }
          }
        });
      });
    });

    context('can search and change page in dataBag', () => {
      it('can search a DataBag and check if empty or not', () => {
      getPolicyFile().then((response) => {
        if (checkResponse(response)) {
          cy.get('[data-cy=search-filter]').type(policyFileName);
          cy.get('[data-cy=search-entity]').click();
          cy.get('[data-cy=policy-file-table-container]').then(body => {
            if (body.text().includes(policyFileName)) {
              cy.get('[data-cy=policy-file-table-container]')
              .contains(policyFileName).should('exist');
            }
          });

          cy.get('[data-cy=search-filter]').clear();
          cy.get('[data-cy=search-entity]').click();
        }
      });
    });
  });

    it('can delete policy file', () => {
      getPolicyFile().then((response) => {
        if (checkResponse(response)) {
          cy.get('[data-cy=search-filter]').type(`${cypressPrefix}-policyFile-${now}`);
          cy.get('[data-cy=search-entity]').click();
            cy.get('[data-cy=policy-file-table-container]').contains(policyFileName)
              .should('exist');
            cy.get('app-policy-files [data-cy=policy-file-table-container] chef-td a')
              .contains(policyFileName).parent().parent().find('.mat-select-trigger').click();

            cy.get('[data-cy=delete]').should('be.visible')
              .click();
            // accept dialog
            cy.get('app-policy-files chef-button').contains('Delete').click();
            // verify success notification and then dismiss it
            cy.get('app-notification.info').
              contains(`Successfully deleted policy file - ${policyFileName}.`);
            cy.get('app-notification.info chef-icon').click();

            cy.get('[data-cy=search-filter]').clear();
            cy.get('[data-cy=search-entity]').click();
          }
        });
    });
  });
});

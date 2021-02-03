import React, { useState } from 'react';
import { Collapse, 
         Navbar, 
         NavbarToggler, 
         NavbarBrand, 
         Nav, NavItem, 
         NavLink 
       } from 'reactstrap';
import { LINK_INCOGNITO_COMMUNITY, 
         LINK_INCOGNITO_DOCUMENTATION, 
         INCOGNITO_FAUCET, 
         COMMUNITY, 
         DOCUMENTATION
       } from './constant'

const Header = (props) => {
  const [collapsed, setCollapsed] = useState(true);
  const toggleNavbar = () => setCollapsed(!collapsed);

  return (
    <div>
      <Navbar color="faded" light>
        <NavbarBrand href="/" className="mr-auto">{INCOGNITO_FAUCET}</NavbarBrand>
        <NavbarToggler onClick={toggleNavbar} className="mr-2" />
        <Collapse isOpen={!collapsed} navbar>
          <Nav navbar>
            <NavItem>
              <NavLink href={LINK_INCOGNITO_COMMUNITY} >{COMMUNITY}</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href={LINK_INCOGNITO_DOCUMENTATION}>{DOCUMENTATION}</NavLink>
            </NavItem>
          </Nav>
        </Collapse>
      </Navbar>
    </div>
  );
}

export default Header;
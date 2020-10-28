import React from 'react';
import {withRouter} from 'react-router-dom';
import {Layout, Menu, Breadcrumb, Dropdown, Avatar, message} from 'antd';
import Icon, {UserOutlined, LaptopOutlined, NotificationOutlined, SettingFilled, UpOutlined} from '@ant-design/icons';
import Logo from './snowflakes.jpg';
import {adminRoutes} from '../../routes';
import './frame.css';
import {clearToken} from "../../utils/auth";

// const {SubMenu} = Menu;
const {Header, Content, Sider} = Layout;
const routes = adminRoutes.filter(route => route.isShow)

function Index(props) {
    const popMenu = (
        <Menu onClick={(p)=>{
            if(p.key === 'logOut'){
                clearToken();
                props.history.push('/login')
            }else{
                message.info(p.key); // tip
            }
        }}>
            <Menu.Item key="note" icon={<NotificationOutlined/>}>通知中心</Menu.Item>
            <Menu.Item key="setting" icon={<SettingFilled/>}>设置</Menu.Item>
            <Menu.Item key="logOut" icon={<LaptopOutlined/>}>退出</Menu.Item>
        </Menu>);
    return (
        <Layout>
            <Header className="header" style={{
                background: "#428bca",
            }}>
                <div className="logo">
                    <img src={Logo} alt='logo' style={{
                        width: '120px',
                        height: "60px"
                    }}/>
                </div>
                <Dropdown overlay={popMenu}>
                    <div>
                        <Avatar>U</Avatar>
                        &nbsp;
                        <UpOutlined/>
                        &nbsp;
                        <span>超级管理员</span>
                    </div>
                </Dropdown>
            </Header>
            <Content style={{padding: '0 50px'}}>
                <Breadcrumb style={{margin: '16px 0'}}>
                    <Breadcrumb.Item>Home</Breadcrumb.Item>
                    <Breadcrumb.Item>List</Breadcrumb.Item>
                    <Breadcrumb.Item>App</Breadcrumb.Item>
                </Breadcrumb>
                <Layout className="site-layout-background" style={{padding: '24px 0'}}>
                    <Sider className="site-layout-background" width={200}>
                        <Menu
                            mode="inline"
                            defaultSelectedKeys={['1']}
                            defaultOpenKeys={['sub1']}
                            style={{height: '100%'}}
                        >
                            {routes.map(route => {
                                return (
                                    <Menu.Item key={route.path} icon={<route.icon/>}
                                               onClick={p => props.history.push(p.key)}>
                                        {route.title}
                                    </Menu.Item>)
                            })}
                        </Menu>
                    </Sider>
                    <Content style={{padding: '0 24px', minHeight: 400}}>{props.children}</Content>
                </Layout>
            </Content>
            {/* <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer> */}
        </Layout>
    )
}

export default withRouter(Index)

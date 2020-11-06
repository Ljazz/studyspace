import React from 'react';
import { Form, Card, Input, Button, message } from 'antd';

function Edit(props) {
    // 验证
    const handleSubmit = e => {
        console.log(e);
        e.preventDefault();
        props.form.validateFieldsAndScroll((err, values) => {
            if (!err) {
                console.log('提交')
            } else {
                message.error('请输入正确的内容')
            }
        });
    };
    return (
        <Card title="商品编辑">
            <Form onSubmit={e => handleSubmit(e)}>
                <Form.Item label="名字" name="name" rules={[{required: true,message: '请输出商品名字'}]}>
                    <Input placeholder="请输入商品名字" />
                </Form.Item>
                <Form.Item label="价格" name='price' rules={[{required: true,message: '请输出商品价格'}]}>
                    <Input placeholder="请输出商品价格" />
                </Form.Item>
                <Form.Item><Button htmlType="submit" type="primary">保存</Button></Form.Item>
            </Form>
        </Card>
    )
}
export default Edit;
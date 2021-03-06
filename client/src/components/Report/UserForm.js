import React from "react";
import { Form, Input, DatePicker, Button, Alert } from "antd";

import styles from "./report.module.scss";

const { TextArea } = Input;

function UserForm({ onFinish, onFinishFailed, error }) {
  return (
    <div>
      <Form
        name="basic"
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        className={styles.form}
      >
        {error ? (
          <Alert
            message="Error While Submitting"
            description="Encountered an error while submitting the form. Please ensure that all the fields are filled / valid and try again"
            type="error"
          />
        ) : (
          ""
        )}

        <h1>Contact Details</h1>

        <Form.Item
          label="Name"
          name="name"
          rules={[
            {
              required: true,
              message: "Please input your Name!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Email ID"
          name="email"
          rules={[
            {
              required: true,
              message: "Please input your Email ID!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Phone No."
          name="phone"
          rules={[
            {
              required: true,
              message: "Please input your Phone No!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <h1>Car Details</h1>

        <Form.Item
          label="Car Model"
          name="model"
          rules={[
            {
              required: true,
              message: "Please input your Car Model!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Car Registration No."
          name="regNo"
          rules={[
            {
              required: true,
              message: "Please input your Car Registration No.!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Car Color"
          name="color"
          rules={[
            {
              required: true,
              message: "Please input your Car Color!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Last Seen Date"
          name="lastSeenDate"
          rules={[
            {
              required: true,
              message: "Please input your Last Seen Date!",
            },
          ]}
        >
          <DatePicker />
        </Form.Item>

        <Form.Item
          label="Last Known Location"
          name="location"
          rules={[
            {
              required: true,
              message: "Please input your Last Known Location!",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Description / Tips"
          name="description"
          rules={[
            {
              required: true,
              message: "Please input some Description!",
            },
          ]}
        >
          <TextArea />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            Create Report
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
}

export default UserForm;

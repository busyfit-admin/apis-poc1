import { Handler } from 'aws-lambda';

export const handler: Handler = async (event, context) => {
    console.log('Hello World');

    console.log('Event:', JSON.stringify(event, null, 2));
    console.log('Context:', JSON.stringify(context, null, 2));

};
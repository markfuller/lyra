import {action, logger, PluginLogger, resource, ServiceBuilder} from 'lyra-workflow';
import * as util from 'util';

import * as Example from './types/Example';

const wf = {
  source: __filename,
  activities: {
    person: resource({
      output: 'name',
      state: (): Example.Person => new Example.Person({
        age: 77,
        name: 'Bert',
        human: false,
      })
    }),
  }
}

const sb = new ServiceBuilder('My::Service');
sb.workflow(wf);
const server = sb.build(global);
logger.info('Starting the server', 'serverId', server.serviceId.toString());
server.start();

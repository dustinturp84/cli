# CloudAMQP CLI

A command line interface for the CloudAMQP API that provides complete management of CloudAMQP instances, VPCs, and instance-specific operations.

## Features

- **Complete API Coverage**: Supports all endpoints from both the main CloudAMQP API and Instance API
- **Smart Key Management**: Automatically handles both main API keys and instance-specific API keys
- **JSON Configuration**: Structured configuration with automatic legacy migration
- **User-Friendly**: Clear help messages, examples, and safety confirmations
- **Error Handling**: Proper API error extraction and display

## Installation

### Build from Source

```bash
go mod download
go build -o cloudamqp
```

### Usage

```bash
./cloudamqp --help
```

## Configuration

### API Key Configuration

The CLI looks for your API key in the following order:

1. `CLOUDAMQP_APIKEY` environment variable
2. `~/.cloudamqprc` file (JSON format)
3. If neither exists, you will be prompted to enter it

Instance API keys are automatically saved when using the `instance get` command.

### Config File Format

The configuration file `~/.cloudamqprc` uses JSON format:

```json
{
  "main_api_key": "your-main-api-key-here",
  "instance_keys": {
    "1234": "instance-1234-api-key",
    "5678": "instance-5678-api-key"
  }
}
```

### Environment Variables

- `CLOUDAMQP_APIKEY` - Main API key
- `CLOUDAMQP_INSTANCE_<id>_APIKEY` - Instance-specific API key (e.g., `CLOUDAMQP_INSTANCE_1234_APIKEY`)

## Commands

### Instance Management

Manage CloudAMQP instances using the main API key.

```bash
# Create a new instance
cloudamqp instance create --name=my-instance --plan=bunny-1 --region=amazon-web-services::us-east-1

# List all instances
cloudamqp instance list

# Get instance details (automatically saves instance API key)
cloudamqp instance get 1234

# Update instance properties
cloudamqp instance update 1234 --name=new-name --plan=rabbit-1

# Resize instance disk
cloudamqp instance resize 1234 --disk-size=100 --allow-downtime

# Delete instance (with confirmation)
cloudamqp instance delete 1234
```

### VPC Management

Manage Virtual Private Clouds.

```bash
# Create a VPC
cloudamqp vpc create --name=my-vpc --region=amazon-web-services::us-east-1 --subnet=10.56.72.0/24

# List all VPCs
cloudamqp vpc list

# Get VPC details
cloudamqp vpc get 5678

# Update VPC
cloudamqp vpc update 5678 --name=new-vpc-name

# Delete VPC (with confirmation)
cloudamqp vpc delete 5678
```

### Instance-Specific Management

Manage specific instances using instance API keys. These commands use the Instance API.

#### Node Management

```bash
# List nodes in an instance
cloudamqp instance manage 1234 nodes list

# Get available versions for upgrade
cloudamqp instance manage 1234 nodes versions
```

#### Plugin Management

```bash
# List available RabbitMQ plugins
cloudamqp instance manage 1234 plugins list
```

#### Instance Actions

```bash
# Restart RabbitMQ
cloudamqp instance manage 1234 actions restart-rabbitmq
cloudamqp instance manage 1234 actions restart-rabbitmq --nodes=node1,node2

# Cluster operations
cloudamqp instance manage 1234 actions restart-cluster
cloudamqp instance manage 1234 actions stop-cluster
cloudamqp instance manage 1234 actions start-cluster

# Instance lifecycle
cloudamqp instance manage 1234 actions stop
cloudamqp instance manage 1234 actions start
cloudamqp instance manage 1234 actions reboot

# Management interface
cloudamqp instance manage 1234 actions restart-management

# Upgrades (asynchronous operations)
cloudamqp instance manage 1234 actions upgrade-erlang
cloudamqp instance manage 1234 actions upgrade-rabbitmq --version=3.10.7
cloudamqp instance manage 1234 actions upgrade-all

# Get target upgrade versions
cloudamqp instance manage 1234 actions upgrade-versions

# Toggle features
cloudamqp instance manage 1234 actions toggle-hipe --enable=true
cloudamqp instance manage 1234 actions toggle-firehose --enable=true --vhost=/
```

#### Account Operations

```bash
# Rotate instance password
cloudamqp instance manage 1234 account rotate-password

# Rotate instance API key
cloudamqp instance manage 1234 account rotate-apikey
```

### Informational Commands

```bash
# List available regions
cloudamqp regions
cloudamqp regions --provider=amazon-web-services

# List available plans
cloudamqp plans
cloudamqp plans --backend=rabbitmq
```


### Team Management

```bash
# List team members
cloudamqp team list

# Invite new team member
cloudamqp team invite --email=user@example.com --role=admin --tags=production

# Update team member
cloudamqp team update --user-id=uuid-here --role=devops

# Remove team member
cloudamqp team remove --email=user@example.com
```

### Administrative Commands

```bash
# Export audit log
cloudamqp audit
cloudamqp audit --timestamp=2024-01

# Rotate main API key
cloudamqp rotate-key
```

## Examples

### Complete Workflow

```bash
# 1. Create an instance
cloudamqp instance create --name=production --plan=bunny-1 --region=amazon-web-services::us-east-1

# 2. Get instance details (saves instance API key)
cloudamqp instance get 1234

# 3. Check instance nodes
cloudamqp instance manage 1234 nodes list

# 4. Install plugins (if needed)
cloudamqp instance manage 1234 plugins list

# 5. Restart RabbitMQ
cloudamqp instance manage 1234 actions restart-rabbitmq

# 6. Upgrade when needed
cloudamqp instance manage 1234 actions upgrade-all
```

### Team Setup

```bash
# Invite team members
cloudamqp team invite --email=dev1@company.com --role=devops
cloudamqp team invite --email=dev2@company.com --role=member

# List current team
cloudamqp team list
```

### Monitoring Setup

```bash
# Check available regions and plans
cloudamqp regions --provider=amazon-web-services
cloudamqp plans --backend=rabbitmq

# Create VPC for isolation
cloudamqp vpc create --name=prod-vpc --region=amazon-web-services::us-east-1 --subnet=10.0.0.0/24

# Create instance in VPC
cloudamqp instance create --name=prod-instance --plan=rabbit-1 --region=amazon-web-services::us-east-1 --vpc-id=5678
```

## Error Handling

The CLI provides clear error messages for common issues:

- **401 Unauthorized**: Check your API key configuration
- **404 Not Found**: Verify instance/VPC IDs are correct
- **400 Bad Request**: Check required parameters and formats

When instance API keys are missing, the CLI will guide you to retrieve them:

```
Error: instance API key not found for instance 1234. Use 'cloudamqp instance get 1234' to retrieve it
```

## Advanced Usage

### Using Environment Variables

```bash
# Set main API key
export CLOUDAMQP_APIKEY="your-main-api-key"

# Set instance-specific API key
export CLOUDAMQP_INSTANCE_1234_APIKEY="instance-specific-key"

# Use the CLI without prompts
cloudamqp instance list
```

### Scripting

The CLI is designed for scripting with:

- JSON output for structured data
- Exit codes for success/failure
- `--force` flags to skip confirmations
- Environment variable support

```bash
#!/bin/bash

# Create instance and capture ID
RESULT=$(cloudamqp instance create --name=temp-instance --plan=lemming --region=amazon-web-services::us-east-1)
INSTANCE_ID=$(echo "$RESULT" | jq -r '.id')

# Get instance details to save API key
cloudamqp instance get "$INSTANCE_ID"

# Perform operations
cloudamqp instance manage "$INSTANCE_ID" actions restart-rabbitmq

# Cleanup
cloudamqp instance delete "$INSTANCE_ID" --force
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and questions:

1. Check the CLI help: `cloudamqp --help`
2. Verify your API key configuration
3. Check the CloudAMQP API documentation
4. Create an issue in this repository

## API Documentation

- [Main CloudAMQP API](https://docs.cloudamqp.com/api.html)
- [Instance API](https://docs.cloudamqp.com/instance-api.html)
- [Terraform Provider](https://registry.terraform.io/providers/cloudamqp/cloudamqp/latest/docs)
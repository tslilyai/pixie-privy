#pragma once

#include <memory>
#include <string>
#include <unordered_map>

#include "src/common/base/base.h"
#include "src/stirling/proto/collector_config.pb.h"
#include "src/stirling/source_registry.h"

namespace pl {
namespace stirling {

/**
 * @brief Convenience function to subscribe to all info classes of
 * a published proto message. This should actually be in an agent.
 * TODO(kgandhi): Move to agent or common utils for agent when available.
 *
 * @param publish_proto
 * @return stirlingpb::Subscribe
 */
stirlingpb::Subscribe SubscribeToAllInfoClasses(const stirlingpb::Publish& publish_proto);

/**
 * The data collector collects data from various different 'sources',
 * and makes them available via a structured API, where the data can then be used and queried as
 * needed (by Pixie or others). Its function is to unify various, disparate sources of data into a
 * common, structured data format.
 */
class Stirling : public NotCopyable {
 public:
  Stirling() = default;
  virtual ~Stirling() = default;

  /**
   * @brief Create a Stirling object
   * Factory method to create Stirling with a default registry containing
   * all sources
   *
   * @return std::unique_ptr<Stirling>
   */
  static std::unique_ptr<Stirling> Create();

  /**
   * @brief Create a Stirling object
   * Factory method to create Stirling with a source registry.
   *
   * @param registry
   * @return std::unique_ptr<Stirling>
   */
  static std::unique_ptr<Stirling> Create(std::unique_ptr<SourceRegistry> registry);

  /**
   * @brief Populate the Publish Proto object. Agent calls this function to get the Publish
   * proto message. The proto publish message contains information (InfoClassSchema) on
   * all the Source Connectors that can be run to gather data and information on the types
   * for the data. The agent can then subscribe to a subset of the published message. The proto
   * is defined in //src/stirling/proto/collector_config.proto.
   *
   */
  virtual void GetPublishProto(stirlingpb::Publish* publish_pb) = 0;

  /**
   * @brief Get the Subscription object. Receive a Subscribe proto message from the agent.
   * Update the schemas based on the subscription message. Generate the appropriate tables
   * that conform to subscription information.
   *
   * @param subscribe_proto
   * @return Status
   */
  virtual Status SetSubscription(const stirlingpb::Subscribe& subscribe_proto) = 0;

  /**
   * @brief Register call-back from Agent. Used to periodically send data.
   *
   * Function signature is:
   *   uint64_t table_id
   *   std::unique_ptr<ColumnWrapperRecordBatch> data
   */
  virtual void RegisterCallback(PushDataCallback f) = 0;

  /**
   * @brief Return a map of table ID to InfoClassManager names.
   */
  virtual std::unordered_map<uint64_t, std::string> TableIDToNameMap() = 0;

  /**
   * @brief Main data collection call. This version blocks, so make sure to wrap a thread around it.
   */
  virtual void Run() = 0;

  /**
   * @brief Main data collection call. This version spawns off as an independent thread.
   */
  virtual Status RunAsThread() = 0;

  /**
   * @brief Stop the running thread. Return will not be immediate.
   */
  virtual void Stop() = 0;

  /**
   * @brief Wait for the running thread to terminate. Assuming you ran RunThread().
   */
  virtual void WaitForThreadJoin() = 0;
};

}  // namespace stirling
}  // namespace pl

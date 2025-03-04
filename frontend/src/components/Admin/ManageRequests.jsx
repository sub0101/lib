import React, { useState } from 'react';
import style from '../../style/managerequest.module.css';
import { useQuery } from '@tanstack/react-query';
import { getAllAdminRequest, getAllRequest } from '../../react-queries/api/requestTracking';

const ManageRequests = () => {
    //   const [requests, setRequests] = useState([
    //     { id: 1, book: "The Great Gatsby", type: "Borrow", date: "2025-03-04" },
    //     { id: 2, book: "1984", type: "Return", date: "2025-03-05" },
    //     { id: 3, book: "To Kill a Mockingbird", type: "Borrow", date: "2025-03-06" },
    //   ]);

    const { data: requests } = useQuery({
        queryFn: getAllAdminRequest,
        queryKey: ["getRequests"]
    })

    const handleApprove = (id) => {
        alert(`Request with ID ${id} approved!`);
        // Logic to approve the request
    };

    const handleReject = (id) => {
        alert(`Request with ID ${id} rejected!`);
        // Logic to reject the request
    };

    return (
        <div className={style.adminContainer}>
            <h1 className={style.adminTitle}>Admin - Manage Requests</h1>
            <table className={style.requestTable}>
                <thead>
                    <tr>
                        <th>Request Book</th>
                        <th>Request ID</th>
                        <th>Request Type</th>
                        <th>Request Date</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {requests && requests.map((request) => (
                        <tr key={request.req_id}>
                            <td>{request.title}</td>
                            <td>{request.req_id}</td>
                            <td>{request.request_type}</td>
                            <td>{request.request_date}</td>
                            {!request.approver_id ?
                                <td className={style.actionButtons}>


                                    <button
                                        className={style.approveButton}
                                        onClick={() => handleApprove(request.id)}>
                                        Approve
                                    </button>

                                    <button
                                        className={style.rejectButton}
                                        onClick={() => handleReject(request.id)}>
                                        Reject
                                    </button>




                                </td>
                                :
                                <td className={style.actionButtons}>
                                    <button
                                        className={style.approveButton}
                                        onClick={() => handleApprove(request.id)}>
                                        {request.request_status}
                                    </button>
                                </td>
                            }
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default ManageRequests;

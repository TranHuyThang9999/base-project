'use client';
import React, { useState, useEffect } from 'react';
import {message} from "antd";

const priceCar = 30000; // Update prices in VND
const priceMotoBike = 5000; // Update prices in VND

// Helper function to format currency values in VND
const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(amount);
};

// Helper function to format date and time
const formatDateAndTime = () => {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0'); // Months are 0-based
    const day = String(now.getDate()).padStart(2, '0');
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
};

export default function ManagerRices() {
    const [list, setList] = useState<{ type: string; payment: string; price: number; date: string; quantity: number }[]>(
        [],
    );
    const [messageApi, contextHolder] = message.useMessage();

    const [filteredList, setFilteredList] = useState(list);
    const [vehicleType, setVehicleType] = useState('Car');
    const [paymentType, setPaymentType] = useState('Cash');
    const [quantity, setQuantity] = useState(1); // Default quantity is 1
    const [filterDate, setFilterDate] = useState('');
    const [currentPage, setCurrentPage] = useState(1);
    const [modalVisible, setModalVisible] = useState(false);

    const itemsPerPage = 5; // Number of items to show per page

    // Load data from localStorage initially
    useEffect(() => {
        const storedData = JSON.parse(localStorage.getItem('data') || '[]');
        setList(storedData);
    }, []);

    // Save `list` to localStorage whenever it changes
    useEffect(() => {
        localStorage.setItem('data', JSON.stringify(list));
    }, [list]);

    // Filter the list by date when `filterDate` changes
    useEffect(() => {
        if (filterDate) {
            const filtered = list.filter((item) => item.date.startsWith(filterDate));
            setFilteredList(filtered);
        } else {
            setFilteredList(list);
        }
    }, [list, filterDate]);

    const handleAddTransaction = (e: React.FormEvent) => {
        e.preventDefault();
        const vehiclePrice = vehicleType === 'Car' ? priceCar : priceMotoBike;
        const totalTransactionPrice = vehiclePrice * quantity; // Calculate total price based on quantity
        const newTransaction = {
            type: vehicleType,
            payment: paymentType,
            price: totalTransactionPrice,
            date: formatDateAndTime(),
            quantity, // Include the quantity
        };
        messageApi.success('Thêm giao dịch thành công')
        setList((prev) => [...prev, newTransaction]);
    };

    const handleDeleteTransaction = (index: number) => {
        const newList = [...list];
        newList.splice(index, 1);
        setList(newList);
    };



    // Calculate statistics based on the filtered list
    // Calculate statistics based on the filtered list
    const totalAmount = filteredList.reduce((sum, item) => sum + item.price, 0);
    const totalCash = filteredList
        .filter((item) => item.payment === 'Cash')
        .reduce((sum, item) => sum + item.price, 0);
    const totalBankTransfer = filteredList
        .filter((item) => item.payment === 'Bank Transfer')
        .reduce((sum, item) => sum + item.price, 0);
    const totalCars = filteredList.filter((item) => item.type === 'Car').reduce((sum, item) => sum + item.quantity, 0);
    const totalMotorcycles = filteredList
        .filter((item) => item.type === 'Motorcycle')
        .reduce((sum, item) => sum + item.quantity, 0);

    // Pagination Logic
    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const currentItems = filteredList.slice(startIndex, endIndex);
    const totalPages = Math.ceil(filteredList.length / itemsPerPage);

    const handlePageChange = (next: boolean) => {
        setCurrentPage((prev) => {
            if (next) {
                return prev === totalPages ? prev : prev + 1;
            } else {
                return prev === 1 ? prev : prev - 1;
            }
        });
    };

    return (
        <div className="p-4 max-w-3xl mx-auto flex flex-col h-screen">
            {contextHolder}
            {/* Summary Section */}
            <div className="mb-4 grid grid-cols-1 sm:grid-cols-2 gap-4 bg-blue-100 p-4 rounded-lg">
                <div className="font-bold flex flex-col items-start space-y-2">
                    <p>
                        Tổng Tiền : <span className="text-green-600">{formatCurrency(totalAmount)}</span>
                    </p>
                    <p>
                        Tiền Mặt : <span className="text-blue-600">{formatCurrency(totalCash)}</span>
                    </p>
                    <p>
                        Chuyển Khoản : <span className="text-orange-600">{formatCurrency(totalBankTransfer)}</span>
                    </p>
                    <p>
                        Số Lượng Ô Tô : <span className="text-blue-600">{totalCars}</span>
                    </p>
                    <p>
                        Số Lượng Xe Máy : <span className="text-orange-600">{totalMotorcycles}</span>
                    </p>
                </div>
            </div>

            {/* Filter by Date */}
            <div className="mb-4 flex flex-col sm:flex-row items-center gap-2">
                <label className="text-gray-700 font-medium">Lọc Theo Ngày:</label>
                <input
                    type="date"
                    value={filterDate}
                    onChange={(e) => setFilterDate(e.target.value)}
                    className="border border-gray-300 rounded-lg p-2 w-full sm:w-auto"
                />
                <button
                    onClick={() => setFilterDate('')}
                    className="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 w-full sm:w-auto"
                >
                    Xoá Bộ Lọc
                </button>
            </div>

            {/* Add Transaction Form */}
            <form onSubmit={handleAddTransaction} className="p-4 bg-gray-100 rounded-lg shadow-md space-y-4">
                <h2 className="text-lg font-bold text-gray-800">Thêm Giao Dịch</h2>
                <div>
                    <label className="block text-gray-700 font-medium mb-2">Loại Xe</label>
                    <select
                        value={vehicleType}
                        onChange={(e) => setVehicleType(e.target.value)}
                        className="border border-gray-300 rounded-lg p-2 w-full"
                    >
                        <option value="Car">Ô Tô</option>
                        <option value="Motorcycle">Xe Máy</option>
                    </select>
                </div>

                <div>
                    <label className="block text-gray-700 font-medium mb-2">Hình Thức Thanh Toán</label>
                    <select
                        value={paymentType}
                        onChange={(e) => setPaymentType(e.target.value)}
                        className="border border-gray-300 rounded-lg p-2 w-full"
                    >
                        <option value="Cash">Tiền Mặt</option>
                        <option value="Bank Transfer">Chuyển Khoản</option>
                    </select>
                </div>

                <div>
                    <label className="block text-gray-700 font-medium mb-2">Số Lượng</label>
                    <input
                        type="number"
                        value={quantity}
                        onChange={(e) => setQuantity(Math.max(1, Number(e.target.value)))} // Prevent negative or zero quantities
                        className="border border-gray-300 rounded-lg p-2 w-full"
                        min="1"
                    />
                </div>

                <button
                    type="submit"
                    className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 w-full"
                >
                    Thêm Giao Dịch
                </button>
            </form>

            {/* Show Modal Button */}
            <button
                onClick={() => setModalVisible(true)}
                className="mt-4 px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 w-full"
            >
                Hiển Thị Danh Sách Giao Dịch
            </button>

            {/* Modal */}
            {modalVisible && (
                <div
                    className="fixed inset-0 bg-gray-800 bg-opacity-75 flex justify-center items-center z-50"
                    onClick={() => setModalVisible(false)}
                >
                    <div
                        className="bg-white rounded-lg w-full max-w-2xl p-6 overflow-y-auto max-h-screen"
                        onClick={(e) => e.stopPropagation()}
                    >
                        <h2 className="text-lg font-bold text-gray-800 mb-4">Danh Sách Giao Dịch</h2>
                        <div className="space-y-4">
                            {currentItems.map((item, index) => (
                                <div
                                    key={index}
                                    className="flex flex-col sm:flex-row justify-between items-start sm:items-center bg-gray-100 p-4 rounded-lg shadow-sm space-y-2 sm:space-y-0"
                                >
									<span>{`${item.date} - ${item.type} x${item.quantity} - ${item.payment} - ${formatCurrency(
                                        item.price,
                                    )}`}</span>
                                    <button
                                        onClick={() => handleDeleteTransaction(index)}
                                        className="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 w-full sm:w-auto"
                                    >
                                        Xóa
                                    </button>
                                </div>
                            ))}
                        </div>

                        {/* Pagination Controls */}
                        <div className="mt-6 flex justify-between">
                            <button
                                disabled={currentPage === 1}
                                onClick={() => handlePageChange(false)}
                                className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
                            >
                                Trang Trước
                            </button>
                            <button
                                disabled={currentPage === totalPages}
                                onClick={() => handlePageChange(true)}
                                className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
                            >
                                Trang Tiếp Theo
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
}